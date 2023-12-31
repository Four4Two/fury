package testutil

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/four4two/fury/tests/e2e/contracts/greeter"
	"github.com/four4two/fury/tests/util"
)

// InitFuryEvmData is run after the chain is running, but before the tests are run.
// It is used to initialize some EVM state, such as deploying contracts.
func (suite *E2eTestSuite) InitFuryEvmData() {
	whale := suite.Fury.GetAccount(FundedAccountName)

	// ensure funded account has nonzero erc20 balance
	balance := suite.Fury.GetErc20Balance(suite.DeployedErc20Address, whale.EvmAddress)
	if balance.Cmp(big.NewInt(0)) != 1 {
		panic(fmt.Sprintf("expected funded account (%s) to have erc20 balance", whale.EvmAddress.Hex()))
	}

	// deploy an example contract
	greeterAddr, _, _, err := greeter.DeployGreeter(
		whale.evmSigner.Auth,
		whale.evmSigner.EvmClient,
		"what's up!",
	)
	suite.NoError(err, "failed to deploy a contract to the EVM")
	suite.Fury.ContractAddrs["greeter"] = greeterAddr
}

func (suite *E2eTestSuite) FundFuryErc20Balance(toAddress common.Address, amount *big.Int) EvmTxResponse {
	// funded account should have erc20 balance
	whale := suite.Fury.GetAccount(FundedAccountName)

	data := util.BuildErc20TransferCallData(toAddress, amount)
	nonce, err := suite.Fury.EvmClient.PendingNonceAt(context.Background(), whale.EvmAddress)
	suite.NoError(err)

	req := util.EvmTxRequest{
		Tx:   ethtypes.NewTransaction(nonce, suite.DeployedErc20Address, big.NewInt(0), 1e5, big.NewInt(1e10), data),
		Data: fmt.Sprintf("fund %s with ERC20 balance (%s)", toAddress.Hex(), amount.String()),
	}

	return whale.SignAndBroadcastEvmTx(req)
}
