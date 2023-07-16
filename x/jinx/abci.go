package jinx

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/four4two/fury/x/jinx/keeper"
)

// BeginBlocker updates interest rates
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.ApplyInterestRateUpdates(ctx)
}
