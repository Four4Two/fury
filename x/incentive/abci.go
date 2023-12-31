package incentive

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/four4two/fury/x/incentive/keeper"
)

// BeginBlocker runs at the start of every block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)

	for _, rp := range params.USDFMintingRewardPeriods {
		k.AccumulateUSDFMintingRewards(ctx, rp)
	}
	for _, rp := range params.JinxSupplyRewardPeriods {
		k.AccumulateJinxSupplyRewards(ctx, rp)
	}
	for _, rp := range params.JinxBorrowRewardPeriods {
		k.AccumulateJinxBorrowRewards(ctx, rp)
	}
	for _, rp := range params.DelegatorRewardPeriods {
		k.AccumulateDelegatorRewards(ctx, rp)
	}
	for _, rp := range params.SwapRewardPeriods {
		k.AccumulateSwapRewards(ctx, rp)
	}
	for _, rp := range params.SavingsRewardPeriods {
		k.AccumulateSavingsRewards(ctx, rp)
	}
	for _, rp := range params.EarnRewardPeriods {
		if err := k.AccumulateEarnRewards(ctx, rp); err != nil {
			panic(fmt.Sprintf("failed to accumulate earn rewards: %s", err))
		}
	}
}
