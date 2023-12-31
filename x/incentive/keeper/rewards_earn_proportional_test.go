package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/four4two/fury/x/incentive/keeper"
	"github.com/four4two/fury/x/incentive/types"
	"github.com/stretchr/testify/require"
)

func TestGetProportionalRewardPeriod(t *testing.T) {
	tests := []struct {
		name                  string
		giveRewardPeriod      types.MultiRewardPeriod
		giveTotalBfurySupply  sdkmath.Int
		giveSingleBfurySupply sdkmath.Int
		wantRewardsPerSecond  sdk.DecCoins
	}{
		{
			"full amount",
			types.NewMultiRewardPeriod(
				true,
				"",
				time.Time{},
				time.Time{},
				cs(c("ufury", 100), c("jinx", 200)),
			),
			i(100),
			i(100),
			toDcs(c("ufury", 100), c("jinx", 200)),
		},
		{
			"3/4 amount",
			types.NewMultiRewardPeriod(
				true,
				"",
				time.Time{},
				time.Time{},
				cs(c("ufury", 100), c("jinx", 200)),
			),
			i(10_000000),
			i(7_500000),
			toDcs(c("ufury", 75), c("jinx", 150)),
		},
		{
			"half amount",
			types.NewMultiRewardPeriod(
				true,
				"",
				time.Time{},
				time.Time{},
				cs(c("ufury", 100), c("jinx", 200)),
			),
			i(100),
			i(50),
			toDcs(c("ufury", 50), c("jinx", 100)),
		},
		{
			"under 1 unit",
			types.NewMultiRewardPeriod(
				true,
				"",
				time.Time{},
				time.Time{},
				cs(c("ufury", 100), c("jinx", 200)),
			),
			i(1000), // total bfury
			i(1),    // bfury supply of this specific vault
			dcs(dc("ufury", "0.1"), dc("jinx", "0.2")), // rewards per second rounded to 0 if under 1ufury/1jinx
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rewardsPerSecond := keeper.GetProportionalRewardsPerSecond(
				tt.giveRewardPeriod,
				tt.giveTotalBfurySupply,
				tt.giveSingleBfurySupply,
			)

			require.Equal(t, tt.wantRewardsPerSecond, rewardsPerSecond)
		})
	}
}
