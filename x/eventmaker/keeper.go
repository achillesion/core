package eventmaker

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	ticType "github.com/marbar3778/tic_mark/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper to house the events and tickets
type Keeper struct {
	coinKeeper     bank.Keeper
	eventKey       sdk.StoreKey // store for upcoming and ongoing events
	closedEventKey sdk.StoreKey // store for events that have passed
	cdc            *codec.Codec
}

// NewKeeper : Generate a new keeper when called
func NewKeeper(coinKeeper bank.Keeper, eventKey sdk.StoreKey, closedEventKey sdk.StoreKey, ticketStoreKey sdk.StoreKey, ban, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper:     coinKeeper,
		eventKey:       eventKey,
		closedEventKey: closedEventKey,
		cdc:            cdc,
	}
}

// GETTERS

// GetEvent - Get specific Event
func (k Keeper) GetEvent(ctx sdk.Context, eventName string, storekey sdk.StoreKey) ticType.Event {
	store := ctx.KVStore(storekey)
	event := store.Get([]byte(eventName))
	var Event ticType.Event
	k.cdc.MustUnmarshalBinaryBare(event, &Event)
	return Event
}

// GetEventOwner - Get the owner of the event
func (k Keeper) GetEventOwner(ctx sdk.Context, eventName string) sdk.AccAddress {
	eventData := k.GetEvent(ctx, eventName, k.eventKey)
	return eventData.EventOwnerAddress
}

// GetAllEvents - Get all eventNames from either store not Both
func (k Keeper) GetAllEvents(ctx sdk.Context, storeKey sdk.StoreKey) sdk.Iterator {
	store := ctx.KVStore(storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// get all my events

// SETTERS

// SetEvent - Set event into store
func (k Keeper) SetEvent(ctx sdk.Context, eventName string, eventData ticType.Event,
	storeKey sdk.StoreKey) {
	store := ctx.KVStore(k.eventKey)
	store.Set([]byte(eventName), k.cdc.MustMarshalBinaryBare(eventData))
}

// DeleteEvent - Delete a event from a store
func (k Keeper) DeleteEvent(ctx sdk.Context, eventName string, storeKey sdk.StoreKey) {
	store := ctx.KVStore(storeKey)
	store.Delete([]byte(eventName))
}

// CloseEvent - Take event from actice events and place in inactive event store
func (k Keeper) CloseEvent(ctx sdk.Context, eventName string) {
	eventData := k.GetEvent(ctx, eventName, k.eventKey)
	k.SetEvent(ctx, eventName, eventData, k.closedEventKey)
	k.DeleteEvent(ctx, eventName, k.eventKey)
}

// SetNewOwner - Change Event Owner
func (k Keeper) NewOwner(ctx sdk.Context, eventName string, previousOwnerAddress sdk.Address, newOwnerAddress sdk.AccAddress, newOwner string) {
	eventData := k.GetEvent(ctx, eventName, k.eventKey)
	if eventData.EventOwnerAddress.Equals(previousOwnerAddress) {
		eventData.EventOwner = newOwner
		eventData.EventOwnerAddress = newOwnerAddress
		k.SetEvent(ctx, eventName, eventData, k.eventKey)
	}
}

// CreateEvent - Create event
func (k Keeper) CreateEvent(ctx sdk.Context, eventName string, totalTickets int,
	eventOwner string, eventOwnerAddress sdk.AccAddress, resale bool,
	ticketData ticType.TicketData, eventDetails ticType.EventDetails) {
	eventData := ticType.CreateEvent(eventName, totalTickets, eventOwner,
		eventOwnerAddress, resale, ticketData,
		eventDetails)
	k.SetEvent(ctx, eventName, eventData, k.eventKey)
}
