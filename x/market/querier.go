package market

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryTickets           = "tickets"
	QueryTicket            = "ticket"
	QueryTicketMarketPlace = "ticket_marketplace"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryTickets:
			return queryAllTickets(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown query")
		}
	}
}

func queryAllTickets(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err sdk.Error) {
	var list []string

	event := path[0]

	iterator := k.GetTickets(ctx, event)

	for ; iterator.Valid(); iterator.Next() {
		eventID := string(iterator.Key())
		list = append(list, eventID)
	}

	bz, ok := codec.MarshalJSONIndent(k.cdc, list)
	if ok != nil {
		panic(ok)
	}
	return bz, nil
}
