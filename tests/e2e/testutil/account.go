package testutil

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/go-bip39"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	emtests "github.com/evmos/ethermint/tests"
	emtypes "github.com/evmos/ethermint/types"
	"github.com/stretchr/testify/require"

	"github.com/four4two/fury/app"
	"github.com/four4two/fury/tests/util"
)

type SigningAccount struct {
	name     string
	mnemonic string

	evmPrivKey *ethsecp256k1.PrivKey
	evmSigner  *util.EvmSigner
	evmReqChan chan<- util.EvmTxRequest
	evmResChan <-chan util.EvmTxResponse

	furySigner *util.FurySigner
	sdkReqChan chan<- util.FuryMsgRequest
	sdkResChan <-chan util.FuryMsgResponse

	EvmAuth *bind.TransactOpts

	EvmAddress common.Address
	SdkAddress sdk.AccAddress

	l *log.Logger
}

// GetAccount returns the account with the given name or fails.
func (chain *Chain) GetAccount(name string) *SigningAccount {
	acc, found := chain.accounts[name]
	if !found {
		chain.t.Fatalf("failed to find account with name %s", name)
	}
	return acc
}

// AddNewSigningAccount sets up a new account with a signer for SDK and EVM transactions.
func (chain *Chain) AddNewSigningAccount(name string, hdPath *hd.BIP44Params, chainId, mnemonic string) *SigningAccount {
	if _, found := chain.accounts[name]; found {
		chain.t.Fatalf("account with name %s already exists", name)
	}

	// Fury signing account for SDK side
	privKeyBytes, err := hd.Secp256k1.Derive()(mnemonic, "", hdPath.String())
	require.NoErrorf(chain.t, err, "failed to derive private key from mnemonic for %s: %s", name, err)
	privKey := &ethsecp256k1.PrivKey{Key: privKeyBytes}

	furySigner := util.NewFurySigner(
		chainId,
		chain.EncodingConfig,
		chain.Auth,
		chain.Tx,
		privKey,
		100,
	)

	sdkReqChan := make(chan util.FuryMsgRequest)
	sdkResChan, err := furySigner.Run(sdkReqChan)
	require.NoErrorf(chain.t, err, "failed to start signer for account %s: %s", name, err)

	// Fury signing account for EVM side
	evmChainId, err := emtypes.ParseChainID(chainId)
	require.NoErrorf(chain.t, err, "unable to parse ethermint-compatible chain id from %s", chainId)
	ecdsaPrivKey, err := crypto.HexToECDSA(hex.EncodeToString(privKeyBytes))
	require.NoError(chain.t, err, "failed to generate ECDSA private key from bytes")

	evmSigner, err := util.NewEvmSigner(
		chain.EvmClient,
		ecdsaPrivKey,
		evmChainId,
	)
	require.NoErrorf(chain.t, err, "failed to create evm signer")

	evmReqChan := make(chan util.EvmTxRequest)
	evmResChan := evmSigner.Run(evmReqChan)

	logger := log.New(os.Stdout, fmt.Sprintf("[%s] ", name), log.LstdFlags)

	chain.accounts[name] = &SigningAccount{
		name:     name,
		mnemonic: mnemonic,
		l:        logger,

		evmPrivKey: privKey,
		evmSigner:  evmSigner,
		evmReqChan: evmReqChan,
		evmResChan: evmResChan,

		furySigner: furySigner,
		sdkReqChan: sdkReqChan,
		sdkResChan: sdkResChan,

		EvmAuth: evmSigner.Auth,

		EvmAddress: evmSigner.Address(),
		SdkAddress: furySigner.Address(),
	}

	return chain.accounts[name]
}

// SignAndBroadcastFuryTx sends a request to the signer and awaits its response.
func (a *SigningAccount) SignAndBroadcastFuryTx(req util.FuryMsgRequest) util.FuryMsgResponse {
	a.l.Printf("broadcasting sdk tx. has data = %+v\n", req.Data)
	// send the request to signer
	a.sdkReqChan <- req

	// TODO: timeout awaiting the response.
	// block and await response
	// response is not returned until the msg is committed to a block
	res := <-a.sdkResChan

	// error will be set if response is not Code 0 (success) or Code 19 (already in mempool)
	if res.Err != nil {
		a.l.Printf("response code: %d error: %s\n", res.Result.Code, res.Result.RawLog)
	} else {
		a.l.Printf("response code: %d, hash %s\n", res.Result.Code, res.Result.TxHash)
	}

	return res
}

// EvmTxResponse is util.EvmTxResponse that also includes the Receipt, if available
type EvmTxResponse struct {
	util.EvmTxResponse
	Receipt *ethtypes.Receipt
}

// SignAndBroadcastEvmTx sends a request to the signer and awaits its response.
func (a *SigningAccount) SignAndBroadcastEvmTx(req util.EvmTxRequest) EvmTxResponse {
	a.l.Printf("broadcasting evm tx %+v\n", req.Data)
	// send the request to signer
	a.evmReqChan <- req

	// block and await response
	// response occurs once tx is submitted to pending tx pool.
	// poll for the receipt to wait for it to be included in a block
	res := <-a.evmResChan
	response := EvmTxResponse{
		EvmTxResponse: res,
	}
	// if failed during signing or broadcast, there will never be a receipt.
	if res.Err != nil {
		return response
	}

	// if we don't have a tx receipt within a given timeout, fail the request
	response.Receipt, response.Err = util.WaitForEvmTxReceipt(a.evmSigner.EvmClient, res.TxHash, 10*time.Second)

	return response
}

func (a *SigningAccount) SignRawEvmData(msg []byte) ([]byte, types.PubKey, error) {
	keyringSigner := emtests.NewSigner(a.evmPrivKey)
	return keyringSigner.SignByAddress(a.SdkAddress, msg)
}

// NewFundedAccount creates a SigningAccount for a random account & funds the account from the whale.
func (chain *Chain) NewFundedAccount(name string, funds sdk.Coins) *SigningAccount {
	entropy, err := bip39.NewEntropy(128)
	require.NoErrorf(chain.t, err, "failed to generate entropy for account %s: %s", name, err)
	mnemonic, err := bip39.NewMnemonic(entropy)
	require.NoErrorf(chain.t, err, "failed to create new mnemonic for account %s: %s", name, err)

	acc := chain.AddNewSigningAccount(
		name,
		hd.CreateHDPath(app.Bip44CoinType, 0, 0),
		chain.ChainId,
		mnemonic,
	)

	// don't attempt to fund when no funds are desired
	if funds.IsZero() {
		return acc
	}

	whale := chain.GetAccount(FundedAccountName)
	whale.l.Printf("attempting to fund created account (%s=%s)\n", name, acc.SdkAddress.String())
	res := whale.SignAndBroadcastFuryTx(
		util.FuryMsgRequest{
			Msgs: []sdk.Msg{
				banktypes.NewMsgSend(whale.SdkAddress, acc.SdkAddress, funds),
			},
			GasLimit:  2e5,
			FeeAmount: sdk.NewCoins(sdk.NewCoin(chain.StakingDenom, sdkmath.NewInt(75000))),
			Data:      fmt.Sprintf("initial funding of account %s", name),
		},
	)

	require.NoErrorf(chain.t, res.Err, "failed to fund new account %s: %s", name, res.Err)

	whale.l.Printf("successfully funded [%s]\n", name)

	return acc
}

// GetNonce fetches the next nonce / sequence number for the account.
func (a *SigningAccount) NextNonce() (uint64, error) {
	return a.evmSigner.EvmClient.PendingNonceAt(context.Background(), a.EvmAddress)
}
