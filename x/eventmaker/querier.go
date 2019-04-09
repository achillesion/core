package eventmaker

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints
const (
	QueryUpcomingEvent      = "upcoming_event"
	QueryClosedEvent        = "closed_event"
	QueryUpcomingEventNames = "upcoming_event_names"
	QueryClosedEventNames   = "closed_event_names"
)

// NewQuerier : Query handler
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryUpcomingEvent:
			return queryUpcomingEvent(ctx, path[1:], req, keeper)
		case QueryClosedEvent:
			return queryClosedEvent(ctx, path[1:], req, keeper)
		case QueryUpcomingEventNames:
			return queryUpcomingEventNames(ctx, req, keeper)
		case QueryClosedEventNames:
			return queryClosedEventNames(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown query endpoint")
		}
	}
}

func queryUpcomingEvent(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err sdk.Error) {
	event := path[0]

	value := k.GetEvent(ctx, event, k.eKey)

	bz, err2 := codec.MarshalJSONIndent(k.cdc, value)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryClosedEvent(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err sdk.Error) {
	event := path[0]

	value := k.GetEvent(ctx, event, k.ceKey)

	bz, err2 := codec.MarshalJSONIndent(k.cdc, value)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryUpcomingEventNames(ctx sdk.Context, req abci.RequestQuery, k Keeper) (res []byte, err sdk.Error) {
	var list []string

	iterator := k.GetAllEvents(ctx, k.eKey)

	for ; iterator.Valid(); iterator.Next() {
		eventName := string(iterator.Key())
		list = append(list, eventName)
	}

	bz, err2 := codec.MarshalJSONIndent(k.cdc, list)
	if err2 != nil {
		panic("could not marshal result")
	}

	return bz, nil
}

func queryClosedEventNames(ctx sdk.Context, req abci.RequestQuery, k Keeper) (res []byte, err sdk.Error) {
	var list []string

	iterator := k.GetAllEvents(ctx, k.ceKey)

	for ; iterator.Valid(); iterator.Next() {
		eventName := string(iterator.Key())
		list = append(list, eventName)
	}

	bz, err2 := codec.MarshalJSONIndent(k.cdc, list)
	if err2 != nil {
		panic("could not marshal result")
	}

	return bz, nil
}
