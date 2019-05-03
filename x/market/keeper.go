package market

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	emTypes "github.com/marbar3778/tic_mark/types"
	em "github.com/marbar3778/tic_mark/x/eventmaker" // add expected_types so you aren't importing
)

//  Keeper for the market module
type Keeper struct {
	EventKeeper em.BaseKeeper
	cKeeper     bank.Keeper
	eKey        sdk.StoreKey // upcoming event key where the tickets will be held
	mKey        sdk.StoreKey // marketplace key for reselling
	uKey        sdk.StoreKey // store to keep an array of all the users that have tickets
	cdc         *codec.Codec
}

func NewKeeper(cKeeper bank.Keeper, eKey sdk.StoreKey, mKey sdk.StoreKey, uKey sdk.StoreKey, cdc *codec.Codec, eventKeeper em.BaseKeeper) Keeper {
	return Keeper{
		EventKeeper: eventKeeper,
		cKeeper:     cKeeper,
		eKey:        eKey,
		mKey:        mKey,
		uKey:        uKey,
		cdc:         cdc,
	}
}

// Getters

// Get all tickets of an event
func (k Keeper) GetTickets(ctx sdk.Context, eventID string) sdk.Iterator {
	store := ctx.KVStore(k.eKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) GetMarketPlaceTickets(ctx sdk.Context, eventID string) sdk.Iterator {
	store := ctx.KVStore(k.mKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) GetMarketPlaceTicket(ctx sdk.Context, eventID string, ticketID string) emTypes.Ticket {
	store := ctx.KVStore(k.mKey)
	event := store.Get([]byte(eventID))
	var Tickets []emTypes.Ticket
	k.cdc.MustUnmarshalBinaryBare(event, &Tickets)
	for _, t := range Tickets {
		if t.TicketID == ticketID {
			return t
		}
	}
	panic("no ticket found in the marketplace")
}

// Get Individual Ticket
func (k Keeper) GetTicket(ctx sdk.Context, eventID string, ticketID string) emTypes.Ticket {
	store := ctx.KVStore(k.eKey)
	event := store.Get([]byte(eventID))
	var Tickets []emTypes.Ticket
	k.cdc.MustUnmarshalBinaryBare(event, &Tickets)
	for _, t := range Tickets {
		if t.TicketID == ticketID {
			return t
		}
	}
	panic("no ticket found with that ID")
}

// Get all tickets that a user may have
func (k Keeper) GetUserTickets(ctx sdk.Context, userAddress sdk.AccAddress) []emTypes.Ticket {
	store := ctx.KVStore(k.uKey)
	user := store.Get([]byte(userAddress))
	var Tickets []emTypes.Ticket
	k.cdc.MustUnmarshalBinaryBare(user, &Tickets)
	return Tickets
}

// Setters

// SetTicket into stores
func (k Keeper) SetTicket(ctx sdk.Context, storeKey sdk.StoreKey, eventID string, ticketData emTypes.Ticket) {
	store := ctx.KVStore(storeKey)
	store.Set([]byte(eventID), k.cdc.MustMarshalBinaryBare(ticketData))
}

// Create Ticket based off the data from the event
func (k Keeper) CreateTicket(ctx sdk.Context, eventID string, ownerName string, ownerAddress sdk.AccAddress) { // add ticket to UKey and EKey
	event, ok := k.EventKeeper.GetOpenEvent(ctx, eventID)
	if !ok {
		panic("error")
	}
	ticketData := event.TicketData
	markUp := ticketData.MarkUpAllowed
	var maxMarkUp int64
	maxMarkUp = int64(markUp)

	ticket := emTypes.CreateTicket(ownerName, ownerAddress, eventID,
		ticketData.InitialPrice, ticketData.TicketsSold, ticketData.TotalTickets, maxMarkUp,
		ticketData.Resale, ticketData.InitialPrice)
	ticketData.TicketsSold = ticketData.TicketsSold + 1
	k.SetTicket(ctx, k.eKey, eventID, ticket) // set ticket to event store

	ownerKey := ownerAddress.String()
	k.SetTicket(ctx, k.uKey, ownerKey, ticket) // set ticket to event store
}

// Add the ticket to the market store
func (k Keeper) ResaleTicket(ctx sdk.Context, ticketID string, eventID string) {
	ticket := k.GetTicket(ctx, eventID, ticketID)
	k.SetTicket(ctx, k.mKey, eventID, ticket)
}

func (k Keeper) AddTicketToMarket(ctx sdk.Context, ticketID string, eventID string, markUp int) {
	ticket := k.GetTicket(ctx, ticketID, eventID)
	requestedMarkUp := markUp
	var requested int64
	requested = int64(requestedMarkUp)

	ticket.SetNewPrice(requested)
	// need to check if first time in sale to use orignial price
	// if not then get maxmarkupallowed and proposed price to see if its within purposed price
	// put initial price
	k.SetTicket(ctx, k.mKey, eventID, ticket)
}

func (k Keeper) SellTicket(ctx sdk.Context, ticketID string, eventID string,
	price int, newOwnerName string, newOwnerAddress sdk.AccAddress, ownerAddress sdk.AccAddress) {
	ticket := k.GetMarketPlaceTicket(ctx, eventID, ticketID)
	if !ownerAddress.Equals(ticket.OwnerAddress) {
		panic("Incorrect Owner")
	}
	ticket.ResaleTicket(newOwnerName, newOwnerAddress)
	ticketEvent := k.GetTicket(ctx, eventID, ticketID)
	ticketEvent = ticket
	k.SetTicket(ctx, k.eKey, eventID, ticketEvent)

}

// uStore := ctx.KVStore(k.uKey)
// uStore.Delete make it delete a single entry of the key not the key
// mStore := ctx.KVStore(k.mKey)
// mStore.Delete([]byte(ticketID))
