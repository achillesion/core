package eventmaker

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
	emTypes "github.com/marbar3778/tic_mark/types"
)

// Keeper to house the events and tickets
type Keeper struct {
	coinKeeper bank.Keeper
	eKey       sdk.StoreKey // store for upcoming and ongoing events
	ceKey      sdk.StoreKey // store for events that have passed
	cdc        *codec.Codec
}

// NewKeeper : Generate a new keeper when called
func NewKeeper(coinKeeper bank.Keeper, eventKey sdk.StoreKey, closedEventKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		eKey:       eventKey,
		ceKey:      closedEventKey,
		cdc:        cdc,
	}
}

// GETTERS

// GetClosedEvent - Get specific Event
func (k Keeper) GetClosedEvent(ctx sdk.Context, eventID string) (event emTypes.Event, ok bool) {
	store := ctx.KVStore(k.ceKey)
	eventData := store.Get([]byte(eventID))
	if eventData == nil {
		return nil, false
	}
	var Event emTypes.Event
	k.cdc.MustUnmarshalBinaryLengthPrefixed(eventData, &Event)
	return Event, true
}

// Get only open events
func (k Keeper) GetOpenEvent(ctx sdk.Context, eventID string) (event emTypes.Event, ok bool) {
	store := ctx.KVStore(k.eKey)
	eventData := store.Get([]byte(eventID))
	if eventData == nil {
		return nil, false
	}
	var Event emTypes.Event
	k.cdc.MustUnmarshalBinaryBare(eventData, &Event)
	return Event, true
}

// GetEventOwner - Get the owner of the event
func (k Keeper) GetEventOwner(ctx sdk.Context, eventID string) sdk.AccAddress {
	eventData, ok := k.GetOpenEvent(ctx, eventID)
	if !ok {
		panic("Event is not found")
	}
	return eventData.EventOwnerAddress
}

// GetAllEvents - Get all eventNames from either store not Both
func (k Keeper) GetAllEvents(ctx sdk.Context, storeKey sdk.StoreKey) sdk.Iterator {
	store := ctx.KVStore(storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// SETTERS

// SetEvent - Set event into store
func (k Keeper) SetEvent(ctx sdk.Context, eventID string, eventData emTypes.Event,
	storeKey sdk.StoreKey) {
	store := ctx.KVStore(k.eKey)
	store.Set([]byte(eventID), k.cdc.MustMarshalBinaryBare(eventData))
}

// DeleteEvent - Delete a event from a store
func (k Keeper) DeleteEvent(ctx sdk.Context, eventID string, storeKey sdk.StoreKey) {
	store := ctx.KVStore(storeKey)
	store.Delete([]byte(eventID))
}

// CloseEvent - Take event from actice events and place in inactive event store
func (k Keeper) CloseEvent(ctx sdk.Context, eventID string) {
	eventData, ok := k.GetOpenEvent(ctx, eventID)
	if !ok {
		panic("No event to close")
	}
	k.SetEvent(ctx, eventID, eventData, k.ceKey)
	k.DeleteEvent(ctx, eventID, k.eKey)
}

// SetNewOwner - Change Event Owner
func (k Keeper) NewOwner(ctx sdk.Context, eventID string, previousOwnerAddress sdk.Address, newOwnerAddress sdk.AccAddress, newOwner string) {
	eventData, ok := k.GetOpenEvent(ctx, eventID)
	if !ok {
		panic("Event does not exist")
	}
	if eventData.EventOwnerAddress.Equals(previousOwnerAddress) {
		eventData.EventOwner = newOwner
		eventData.EventOwnerAddress = newOwnerAddress
		k.SetEvent(ctx, eventID, eventData, k.eKey)
	}
}

// CreateEvent - Create event
func (k Keeper) CreateEvent(ctx sdk.Context, eventName string, totalTickets int,
	eventOwner string, eventOwnerAddress sdk.AccAddress, resale bool,
	ticketData emTypes.TicketData, eventDetails emTypes.EventDetails) {
	eventData := emTypes.CreateEvent(eventName, totalTickets, eventOwner,
		eventOwnerAddress, resale, ticketData,
		eventDetails)
	k.SetEvent(ctx, eventName, eventData, k.eKey)
}

// Mark ticket as checkedin
