package auction

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/four4two/fury/x/auction/keeper"
	"github.com/four4two/fury/x/auction/types"
)

// BeginBlocker closes all expired auctions at the end of each block. It panics if
// there's an error other than ErrAuctionNotFound.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	err := k.CloseExpiredAuctions(ctx)
	if err != nil && !errors.Is(err, types.ErrAuctionNotFound) {
		panic(err)
	}
}
