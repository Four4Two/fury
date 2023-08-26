package earn_test

import (
	"testing"

	"github.com/four4two/fury/app"
	"github.com/four4two/fury/x/earn"
	"github.com/four4two/fury/x/earn/testutil"
	"github.com/four4two/fury/x/earn/types"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type genesisTestSuite struct {
	testutil.Suite
}

func (suite *genesisTestSuite) Test_InitGenesis_ValidationPanic() {
	invalidState := types.NewGenesisState(
		types.Params{
			AllowedVaults: types.AllowedVaults{
				types.NewAllowedVault(
					"usdf", types.StrategyTypes{types.STRATEGY_TYPE_JINX},
					false,
					nil,
				),
			},
		},
		types.VaultRecords{
			{
				TotalShares: types.VaultShare{
					Denom: "", Amount: sdk.NewDec(1),
				},
			},
		},
		types.VaultShareRecords{},
	)

	suite.Panics(func() {
		earn.InitGenesis(suite.Ctx, suite.Keeper, suite.AccountKeeper, invalidState)
	}, "expected init genesis to panic with invalid state")
}

func (suite *genesisTestSuite) Test_InitAndExportGenesis() {
	depositor_1, err := sdk.AccAddressFromBech32("fury1esagqd83rhqdtpy5sxhklaxgn58k2m3sa9wnu4")
	suite.Require().NoError(err)
	depositor_2, err := sdk.AccAddressFromBech32("fury1mq9qxlhze029lm0frzw2xr6hem8c3k9tu2gu2x")
	suite.Require().NoError(err)

	// slices are sorted by key as stored in the data store, so init and export can be compared with equal
	state := types.NewGenesisState(
		types.Params{
			AllowedVaults: types.AllowedVaults{
				types.NewAllowedVault(
					"usdf",
					types.StrategyTypes{types.STRATEGY_TYPE_JINX},
					false,
					nil,
				),
				types.NewAllowedVault(
					"ufury",
					types.StrategyTypes{types.STRATEGY_TYPE_SAVINGS},
					true,
					[]sdk.AccAddress{suite.AccountKeeper.GetModuleAddress("distribution")},
				),
			},
		},
		types.VaultRecords{
			types.VaultRecord{
				TotalShares: types.NewVaultShare("ufury", sdk.NewDec(3800000)),
			},
			types.VaultRecord{
				TotalShares: types.NewVaultShare("usdf", sdk.NewDec(1000000)),
			},
		},
		types.VaultShareRecords{
			types.VaultShareRecord{
				Depositor: depositor_1,
				Shares: types.NewVaultShares(
					types.NewVaultShare("usdf", sdk.NewDec(500000)),
					types.NewVaultShare("ufury", sdk.NewDec(1900000)),
				),
			},
			types.VaultShareRecord{
				Depositor: depositor_2,
				Shares: types.NewVaultShares(
					types.NewVaultShare("usdf", sdk.NewDec(500000)),
					types.NewVaultShare("ufury", sdk.NewDec(1900000)),
				),
			},
		},
	)

	earn.InitGenesis(suite.Ctx, suite.Keeper, suite.AccountKeeper, state)
	suite.Equal(state.Params, suite.Keeper.GetParams(suite.Ctx))

	vaultRecord1, _ := suite.Keeper.GetVaultRecord(suite.Ctx, "ufury")
	vaultRecord2, _ := suite.Keeper.GetVaultRecord(suite.Ctx, "usdf")
	suite.Equal(state.VaultRecords[0], vaultRecord1)
	suite.Equal(state.VaultRecords[1], vaultRecord2)

	shareRecord1, _ := suite.Keeper.GetVaultShareRecord(suite.Ctx, depositor_1)
	shareRecord2, _ := suite.Keeper.GetVaultShareRecord(suite.Ctx, depositor_2)

	suite.Equal(state.VaultShareRecords[0], shareRecord1)
	suite.Equal(state.VaultShareRecords[1], shareRecord2)

	exportedState := earn.ExportGenesis(suite.Ctx, suite.Keeper)
	suite.Equal(state, exportedState)
}

func (suite *genesisTestSuite) Test_Marshall() {
	depositor_1, err := sdk.AccAddressFromBech32("fury1esagqd83rhqdtpy5sxhklaxgn58k2m3sa9wnu4")
	suite.Require().NoError(err)
	depositor_2, err := sdk.AccAddressFromBech32("fury1mq9qxlhze029lm0frzw2xr6hem8c3k9tu2gu2x")
	suite.Require().NoError(err)

	// slices are sorted by key as stored in the data store, so init and export can be compared with equal
	state := types.NewGenesisState(
		types.Params{
			AllowedVaults: types.AllowedVaults{
				types.NewAllowedVault(
					"usdf",
					types.StrategyTypes{types.STRATEGY_TYPE_JINX},
					false,
					nil,
				),
				types.NewAllowedVault(
					"ufury",
					types.StrategyTypes{types.STRATEGY_TYPE_SAVINGS},
					true,
					[]sdk.AccAddress{suite.AccountKeeper.GetModuleAddress("distribution")},
				),
			},
		},
		types.VaultRecords{
			types.VaultRecord{
				TotalShares: types.NewVaultShare("ufury", sdk.NewDec(3800000)),
			},
			types.VaultRecord{
				TotalShares: types.NewVaultShare("usdf", sdk.NewDec(1000000)),
			},
		},
		types.VaultShareRecords{
			types.VaultShareRecord{
				Depositor: depositor_1,
				Shares: types.NewVaultShares(
					types.NewVaultShare("usdf", sdk.NewDec(500000)),
					types.NewVaultShare("ufury", sdk.NewDec(1900000)),
				),
			},
			types.VaultShareRecord{
				Depositor: depositor_2,
				Shares: types.NewVaultShares(
					types.NewVaultShare("usdf", sdk.NewDec(500000)),
					types.NewVaultShare("ufury", sdk.NewDec(1900000)),
				),
			},
		},
	)

	encodingCfg := app.MakeEncodingConfig()
	cdc := encodingCfg.Marshaler

	bz, err := cdc.Marshal(&state)
	suite.Require().NoError(err, "expected genesis state to marshal without error")

	var decodedState types.GenesisState
	err = cdc.Unmarshal(bz, &decodedState)
	suite.Require().NoError(err, "expected genesis state to unmarshal without error")

	suite.Equal(state, decodedState)
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(genesisTestSuite))
}
