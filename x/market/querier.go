package market

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryTickets     = "tickets"
	QueryTicket      = "ticket"
	QueryMarketPlace = "ticket_marketplace"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryTickets:
			return queryAllTickets(ctx, path[1:], req, keeper)
		case QueryTicket:
			return queryTicket(ctx, path, req, keeper)
		case QueryMarketPlace:
			return queryMarketPlace(ctx, path, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown query")
		}
	}
}

func queryAllTickets(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err sdk.Error) {

	event := path[0]

	// get all the tickets
	tickets, ok := k.GetTickets(ctx, event)
	if !ok {
		panic("something")
	}

	bz, err2 := codec.MarshalJSONIndent(k.cdc, tickets)
	if err2 != nil {
		panic(err2)
	}
	return bz, nil
}

func queryTicket(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err sdk.Error) {
	event := path[0]
	ticket := path[1]

	value, ok := k.GetTicket(ctx, event, ticket)
	if !ok {
		panic("No ticket")
	}

	bz, err2 := codec.MarshalJSONIndent(k.cdc, value)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}
	return bz, nil
}

func queryMarketPlace(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err sdk.Error) {

	event := path[0]

	tickets, ok := k.GetMarketPlaceTickets(ctx, event)
	if !ok {
		panic("Nothing in the market place for you")
	}

	bz, err2 := codec.MarshalJSONIndent(k.cdc, tickets)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}
	return bz, nil
}
