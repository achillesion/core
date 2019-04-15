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
	var list []string

	event := path[0]

	iterator := k.GetTickets(ctx, event)

	for ; iterator.Valid(); iterator.Next() {
		eventID := string(iterator.Key())
		list = append(list, eventID)
	}

	bz, err2 := codec.MarshalJSONIndent(k.cdc, list)
	if err2 != nil {
		panic(err2)
	}
	return bz, nil
}

func queryTicket(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err sdk.Error) {
	event := path[0]
	ticket := path[1]

	value := k.GetTicket(ctx, event, ticket)

	bz, err2 := codec.MarshalJSONIndent(k.cdc, value)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}
	return bz, nil
}

func queryMarketPlace(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err sdk.Error) {
	var list []string

	event := path[0]

	iterator := k.GetMarketPlaceTickets(ctx, event)

	for ; iterator.Valid(); iterator.Next() {
		eventID := string(iterator.Key())
		list = append(list, eventID)
	}
	bz, err2 := codec.MarshalJSONIndent(k.cdc, list)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}
	return bz, nil
}
