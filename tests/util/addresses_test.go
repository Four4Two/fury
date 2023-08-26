package util_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/four4two/fury/app"
	"github.com/four4two/fury/tests/util"
)

func TestAddressConversion(t *testing.T) {
	app.SetSDKConfig()
	bech32Addr := sdk.MustAccAddressFromBech32("fury17d2wax0zhjrrecvaszuyxdf5wcu5a0p4vpmrwu")
	hexAddr := common.HexToAddress("0xF354EE99E2BC863CE19D80B843353476394EBC35")
	require.Equal(t, bech32Addr, util.EvmToSdkAddress(hexAddr))
	require.Equal(t, hexAddr, util.SdkToEvmAddress(bech32Addr))
}
