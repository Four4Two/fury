package keeper_test

import (
	"os"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/four4two/fury/app"
	"github.com/four4two/fury/x/earn/testutil"
	"github.com/four4two/fury/x/earn/types"
	"github.com/stretchr/testify/suite"
)

func TestMain(m *testing.M) {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	os.Exit(m.Run())
}

type depositTestSuite struct {
	testutil.Suite
}

func (suite *depositTestSuite) SetupTest() {
	suite.Suite.SetupTest()
	suite.Keeper.SetParams(suite.Ctx, types.DefaultParams())
}

func TestDepositTestSuite(t *testing.T) {
	suite.Run(t, new(depositTestSuite))
}

func (suite *depositTestSuite) TestDeposit_Balances() {
	vaultDenom := "usdf"
	startBalance := sdk.NewInt64Coin(vaultDenom, 1000)
	depositAmount := sdk.NewInt64Coin(vaultDenom, 100)

	suite.CreateVault(vaultDenom, types.StrategyTypes{types.STRATEGY_TYPE_JINX}, false, nil)

	acc := suite.CreateAccount(sdk.NewCoins(startBalance), 0)

	err := suite.Keeper.Deposit(suite.Ctx, acc.GetAddress(), depositAmount, types.STRATEGY_TYPE_JINX)
	suite.Require().NoError(err)

	suite.AccountBalanceEqual(
		acc.GetAddress(),
		sdk.NewCoins(startBalance.Sub(depositAmount)), // Account decreases by deposit
	)

	suite.VaultTotalValuesEqual(sdk.NewCoins(depositAmount))
	suite.VaultTotalSharesEqual(types.NewVaultShares(
		types.NewVaultShare(depositAmount.Denom, sdk.NewDecFromInt(depositAmount.Amount)),
	))
}

func (suite *depositTestSuite) TestDeposit_Exceed() {
	vaultDenom := "usdf"
	startBalance := sdk.NewInt64Coin(vaultDenom, 1000)
	depositAmount := sdk.NewInt64Coin(vaultDenom, 1001)

	suite.CreateVault(vaultDenom, types.StrategyTypes{types.STRATEGY_TYPE_JINX}, false, nil)

	acc := suite.CreateAccount(sdk.NewCoins(startBalance), 0)

	err := suite.Keeper.Deposit(suite.Ctx, acc.GetAddress(), depositAmount, types.STRATEGY_TYPE_JINX)
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, sdkerrors.ErrInsufficientFunds)

	// No changes in balances

	suite.AccountBalanceEqual(
		acc.GetAddress(),
		sdk.NewCoins(startBalance),
	)

	suite.ModuleAccountBalanceEqual(
		sdk.NewCoins(),
	)
}

func (suite *depositTestSuite) TestDeposit_Zero() {
	vaultDenom := "usdf"
	startBalance := sdk.NewInt64Coin(vaultDenom, 1000)
	depositAmount := sdk.NewInt64Coin(vaultDenom, 0)

	suite.CreateVault(vaultDenom, types.StrategyTypes{types.STRATEGY_TYPE_JINX}, false, nil)

	acc := suite.CreateAccount(sdk.NewCoins(startBalance), 0)

	err := suite.Keeper.Deposit(suite.Ctx, acc.GetAddress(), depositAmount, types.STRATEGY_TYPE_JINX)
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, types.ErrInsufficientAmount)

	// No changes in balances

	suite.AccountBalanceEqual(
		acc.GetAddress(),
		sdk.NewCoins(startBalance),
	)

	suite.ModuleAccountBalanceEqual(
		sdk.NewCoins(),
	)
}

func (suite *depositTestSuite) TestDeposit_InvalidVault() {
	vaultDenom := "usdf"
	startBalance := sdk.NewInt64Coin(vaultDenom, 1000)
	depositAmount := sdk.NewInt64Coin(vaultDenom, 1001)

	// Vault not created -- doesn't exist

	acc := suite.CreateAccount(sdk.NewCoins(startBalance), 0)

	err := suite.Keeper.Deposit(suite.Ctx, acc.GetAddress(), depositAmount, types.STRATEGY_TYPE_JINX)
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, types.ErrInvalidVaultDenom)

	// No changes in balances

	suite.AccountBalanceEqual(
		acc.GetAddress(),
		sdk.NewCoins(startBalance),
	)

	suite.ModuleAccountBalanceEqual(
		sdk.NewCoins(),
	)
}

func (suite *depositTestSuite) TestDeposit_InvalidStrategy() {
	vaultDenom := "usdf"
	startBalance := sdk.NewInt64Coin(vaultDenom, 1000)
	depositAmount := sdk.NewInt64Coin(vaultDenom, 1001)

	suite.CreateVault(vaultDenom, types.StrategyTypes{types.STRATEGY_TYPE_JINX}, false, nil)

	acc := suite.CreateAccount(sdk.NewCoins(startBalance), 0)

	err := suite.Keeper.Deposit(suite.Ctx, acc.GetAddress(), depositAmount, types.STRATEGY_TYPE_SAVINGS)
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, types.ErrInvalidVaultStrategy)
}

func (suite *depositTestSuite) TestDeposit_PrivateVault() {
	vaultDenom := "usdf"
	startBalance := sdk.NewInt64Coin(vaultDenom, 1000)
	depositAmount := sdk.NewInt64Coin(vaultDenom, 100)

	acc1 := suite.CreateAccount(sdk.NewCoins(startBalance), 0)
	acc2 := suite.CreateAccount(sdk.NewCoins(startBalance), 1)

	suite.CreateVault(
		vaultDenom,
		types.StrategyTypes{types.STRATEGY_TYPE_JINX},
		true,
		[]sdk.AccAddress{acc1.GetAddress()},
	)

	err := suite.Keeper.Deposit(suite.Ctx, acc2.GetAddress(), depositAmount, types.STRATEGY_TYPE_JINX)
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, types.ErrAccountDepositNotAllowed, "private vault should not allow deposits from non-allowed addresses")

	err = suite.Keeper.Deposit(suite.Ctx, acc1.GetAddress(), depositAmount, types.STRATEGY_TYPE_JINX)
	suite.Require().NoError(err, "private vault should allow deposits from allowed addresses")
}

func (suite *depositTestSuite) TestDeposit_bFury() {
	vaultDenom := "bfury"
	coinDenom := testutil.TestBfuryDenoms[0]

	startBalance := sdk.NewInt64Coin(coinDenom, 1000)
	depositAmount := sdk.NewInt64Coin(coinDenom, 100)

	acc1 := suite.CreateAccount(sdk.NewCoins(startBalance), 0)

	// vault denom is only "bfury" which has it's own special handler
	suite.CreateVault(
		vaultDenom,
		types.StrategyTypes{types.STRATEGY_TYPE_SAVINGS},
		false,
		[]sdk.AccAddress{},
	)

	err := suite.Keeper.Deposit(suite.Ctx, acc1.GetAddress(), depositAmount, types.STRATEGY_TYPE_SAVINGS)
	suite.Require().NoError(
		err,
		"should be able to deposit bfury derivative denom in bfury vault",
	)
}
