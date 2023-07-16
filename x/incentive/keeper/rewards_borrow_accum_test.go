package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/four4two/fury/x/incentive/types"
)

type AccumulateBorrowRewardsTests struct {
	unitTester
}

func (suite *AccumulateBorrowRewardsTests) storedTimeEquals(denom string, expected time.Time) {
	storedTime, found := suite.keeper.GetPreviousJinxBorrowRewardAccrualTime(suite.ctx, denom)
	suite.True(found)
	suite.Equal(expected, storedTime)
}

func (suite *AccumulateBorrowRewardsTests) storedIndexesEqual(denom string, expected types.RewardIndexes) {
	storedIndexes, found := suite.keeper.GetJinxBorrowRewardIndexes(suite.ctx, denom)
	suite.Equal(found, expected != nil)

	if found {
		suite.Equal(expected, storedIndexes)
	} else {
		// Can't compare Equal for types.RewardIndexes(nil) vs types.RewardIndexes{}
		suite.Empty(storedIndexes)
	}
}

func TestAccumulateBorrowRewards(t *testing.T) {
	suite.Run(t, new(AccumulateBorrowRewardsTests))
}

func (suite *AccumulateBorrowRewardsTests) TestStateUpdatedWhenBlockTimeHasIncreased() {
	denom := "bnb"

	jinxKeeper := newFakeJinxKeeper().addTotalBorrow(c(denom, 1e6), d("1"))
	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, jinxKeeper, nil, nil, nil, nil, nil, nil)

	suite.storeGlobalBorrowIndexes(types.MultiRewardIndexes{
		{
			CollateralType: denom,
			RewardIndexes: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
		},
	})
	previousAccrualTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	suite.keeper.SetPreviousJinxBorrowRewardAccrualTime(suite.ctx, denom, previousAccrualTime)

	newAccrualTime := previousAccrualTime.Add(1 * time.Hour)
	suite.ctx = suite.ctx.WithBlockTime(newAccrualTime)

	period := types.NewMultiRewardPeriod(
		true,
		denom,
		time.Unix(0, 0), // ensure the test is within start and end times
		distantFuture,
		cs(c("jinx", 2000), c("ufury", 1000)), // same denoms as in global indexes
	)

	suite.keeper.AccumulateJinxBorrowRewards(suite.ctx, period)

	// check time and factors

	suite.storedTimeEquals(denom, newAccrualTime)
	suite.storedIndexesEqual(denom, types.RewardIndexes{
		{
			CollateralType: "jinx",
			RewardFactor:   d("7.22"),
		},
		{
			CollateralType: "ufury",
			RewardFactor:   d("3.64"),
		},
	})
}

func (suite *AccumulateBorrowRewardsTests) TestStateUnchangedWhenBlockTimeHasNotIncreased() {
	denom := "bnb"

	jinxKeeper := newFakeJinxKeeper().addTotalBorrow(c(denom, 1e6), d("1"))
	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, jinxKeeper, nil, nil, nil, nil, nil, nil)

	previousIndexes := types.MultiRewardIndexes{
		{
			CollateralType: denom,
			RewardIndexes: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
		},
	}
	suite.storeGlobalBorrowIndexes(previousIndexes)
	previousAccrualTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	suite.keeper.SetPreviousJinxBorrowRewardAccrualTime(suite.ctx, denom, previousAccrualTime)

	suite.ctx = suite.ctx.WithBlockTime(previousAccrualTime)

	period := types.NewMultiRewardPeriod(
		true,
		denom,
		time.Unix(0, 0), // ensure the test is within start and end times
		distantFuture,
		cs(c("jinx", 2000), c("ufury", 1000)), // same denoms as in global indexes
	)

	suite.keeper.AccumulateJinxBorrowRewards(suite.ctx, period)

	// check time and factors

	suite.storedTimeEquals(denom, previousAccrualTime)
	expected, f := previousIndexes.Get(denom)
	suite.True(f)
	suite.storedIndexesEqual(denom, expected)
}

func (suite *AccumulateBorrowRewardsTests) TestNoAccumulationWhenSourceSharesAreZero() {
	denom := "bnb"

	jinxKeeper := newFakeJinxKeeper() // zero total borrows
	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, jinxKeeper, nil, nil, nil, nil, nil, nil)

	previousIndexes := types.MultiRewardIndexes{
		{
			CollateralType: denom,
			RewardIndexes: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
		},
	}
	suite.storeGlobalBorrowIndexes(previousIndexes)
	previousAccrualTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	suite.keeper.SetPreviousJinxBorrowRewardAccrualTime(suite.ctx, denom, previousAccrualTime)

	firstAccrualTime := previousAccrualTime.Add(7 * time.Second)
	suite.ctx = suite.ctx.WithBlockTime(firstAccrualTime)

	period := types.NewMultiRewardPeriod(
		true,
		denom,
		time.Unix(0, 0), // ensure the test is within start and end times
		distantFuture,
		cs(c("jinx", 2000), c("ufury", 1000)), // same denoms as in global indexes
	)

	suite.keeper.AccumulateJinxBorrowRewards(suite.ctx, period)

	// check time and factors

	suite.storedTimeEquals(denom, firstAccrualTime)
	expected, f := previousIndexes.Get(denom)
	suite.True(f)
	suite.storedIndexesEqual(denom, expected)
}

func (suite *AccumulateBorrowRewardsTests) TestStateAddedWhenStateDoesNotExist() {
	denom := "bnb"

	jinxKeeper := newFakeJinxKeeper().addTotalBorrow(c(denom, 1e6), d("1"))
	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, jinxKeeper, nil, nil, nil, nil, nil, nil)

	period := types.NewMultiRewardPeriod(
		true,
		denom,
		time.Unix(0, 0), // ensure the test is within start and end times
		distantFuture,
		cs(c("jinx", 2000), c("ufury", 1000)),
	)

	firstAccrualTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	suite.ctx = suite.ctx.WithBlockTime(firstAccrualTime)

	suite.keeper.AccumulateJinxBorrowRewards(suite.ctx, period)

	// After the first accumulation only the current block time should be stored.
	// The indexes will be empty as no time has passed since the previous block because it didn't exist.
	suite.storedTimeEquals(denom, firstAccrualTime)
	suite.storedIndexesEqual(denom, nil)

	secondAccrualTime := firstAccrualTime.Add(10 * time.Second)
	suite.ctx = suite.ctx.WithBlockTime(secondAccrualTime)

	suite.keeper.AccumulateJinxBorrowRewards(suite.ctx, period)

	// After the second accumulation both current block time and indexes should be stored.
	suite.storedTimeEquals(denom, secondAccrualTime)
	suite.storedIndexesEqual(denom, types.RewardIndexes{
		{
			CollateralType: "jinx",
			RewardFactor:   d("0.02"),
		},
		{
			CollateralType: "ufury",
			RewardFactor:   d("0.01"),
		},
	})
}

func (suite *AccumulateBorrowRewardsTests) TestNoPanicWhenStateDoesNotExist() {
	denom := "bnb"

	jinxKeeper := newFakeJinxKeeper()
	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, jinxKeeper, nil, nil, nil, nil, nil, nil)

	period := types.NewMultiRewardPeriod(
		true,
		denom,
		time.Unix(0, 0), // ensure the test is within start and end times
		distantFuture,
		cs(),
	)

	accrualTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	suite.ctx = suite.ctx.WithBlockTime(accrualTime)

	// Accumulate with no source shares and no rewards per second will result in no increment to the indexes.
	// No increment and no previous indexes stored, results in an updated of nil. Setting this in the state panics.
	// Check there is no panic.
	suite.NotPanics(func() {
		suite.keeper.AccumulateJinxBorrowRewards(suite.ctx, period)
	})

	suite.storedTimeEquals(denom, accrualTime)
	suite.storedIndexesEqual(denom, nil)
}

func (suite *AccumulateBorrowRewardsTests) TestNoAccumulationWhenBeforeStartTime() {
	denom := "bnb"

	jinxKeeper := newFakeJinxKeeper().addTotalBorrow(c(denom, 1e6), d("1"))
	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, jinxKeeper, nil, nil, nil, nil, nil, nil)

	previousIndexes := types.MultiRewardIndexes{
		{
			CollateralType: denom,
			RewardIndexes: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
		},
	}
	suite.storeGlobalBorrowIndexes(previousIndexes)
	previousAccrualTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	suite.keeper.SetPreviousJinxBorrowRewardAccrualTime(suite.ctx, denom, previousAccrualTime)

	firstAccrualTime := previousAccrualTime.Add(10 * time.Second)

	period := types.NewMultiRewardPeriod(
		true,
		denom,
		firstAccrualTime.Add(time.Nanosecond), // start time after accrual time
		distantFuture,
		cs(c("jinx", 2000), c("ufury", 1000)),
	)

	suite.ctx = suite.ctx.WithBlockTime(firstAccrualTime)

	suite.keeper.AccumulateJinxBorrowRewards(suite.ctx, period)

	// The accrual time should be updated, but the indexes unchanged
	suite.storedTimeEquals(denom, firstAccrualTime)
	expectedIndexes, f := previousIndexes.Get(denom)
	suite.True(f)
	suite.storedIndexesEqual(denom, expectedIndexes)
}

func (suite *AccumulateBorrowRewardsTests) TestPanicWhenCurrentTimeLessThanPrevious() {
	denom := "bnb"

	jinxKeeper := newFakeJinxKeeper().addTotalBorrow(c(denom, 1e6), d("1"))
	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, jinxKeeper, nil, nil, nil, nil, nil, nil)

	previousAccrualTime := time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	suite.keeper.SetPreviousJinxBorrowRewardAccrualTime(suite.ctx, denom, previousAccrualTime)

	firstAccrualTime := time.Time{}

	period := types.NewMultiRewardPeriod(
		true,
		denom,
		time.Time{}, // start time after accrual time
		distantFuture,
		cs(c("jinx", 2000), c("ufury", 1000)),
	)

	suite.ctx = suite.ctx.WithBlockTime(firstAccrualTime)

	suite.Panics(func() {
		suite.keeper.AccumulateJinxBorrowRewards(suite.ctx, period)
	})
}
