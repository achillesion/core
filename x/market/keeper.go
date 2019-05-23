package market

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	emTypes "github.com/marbar3778/tic_mark/types"
	em "github.com/marbar3778/tic_mark/x/eventmaker" // add expected_types so you aren't importing
)

const StoreKey = "market"

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
func (k Keeper) GetTickets(ctx sdk.Context, eventID string) (tickets []emTypes.Ticket, ok bool) {
	store := ctx.KVStore(k.eKey)
	event := store.Get([]byte(eventID))
	if event == nil {
		return nil, false
	}

	var Tickets []emTypes.Ticket
	k.cdc.MustUnmarshalBinaryBare(event, &Tickets)
	return Tickets, true
}

func (k Keeper) GetMarketPlaceTickets(ctx sdk.Context, eventID string) (tickets []emTypes.Ticket, ok bool) {
	store := ctx.KVStore(k.mKey)
	event := store.Get([]byte(eventID))
	if event == nil {
		return nil, false
	}
	var Tickets []emTypes.Ticket
	k.cdc.MustUnmarshalBinaryBare(event, &Tickets)
	return Tickets, true
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
func (k Keeper) GetTicket(ctx sdk.Context, eventID string, ticketID string) (ticket emTypes.Ticket, ok bool) {
	store := ctx.KVStore(k.eKey)
	event := store.Get([]byte(eventID))
	var Tickets []emTypes.Ticket
	k.cdc.MustUnmarshalBinaryBare(event, &Tickets)
	for _, t := range Tickets {
		if t.TicketID == ticketID {
			return t, true
		}
	}
	return emTypes.Ticket{}, false
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
func (k Keeper) SetTicket(ctx sdk.Context, storeKey sdk.StoreKey, eventID string, ticketData []emTypes.Ticket) {
	store := ctx.KVStore(storeKey)
	store.Set([]byte(eventID), k.cdc.MustMarshalBinaryBare(ticketData))
}

// Delete the ticket from the marketplace
func (k Keeper) DeleteMarketTicket(ctx sdk.Context, eventID string, ticketID string) {
	tickets, ok := k.GetMarketPlaceTickets(ctx, eventID)
	if !ok {
		panic("Something")
	}
	for _, ticket := range tickets {
		if ticket.TicketID == ticketID {
			// delete me
		}
	}
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

	// Creation of the ticket
	ticket := emTypes.CreateTicket(ownerName, ownerAddress, eventID,
		ticketData.InitialPrice, ticketData.TicketsSold, ticketData.TotalTickets, maxMarkUp,
		ticketData.Resale, ticketData.InitialPrice)
	ticketData.TicketsSold = ticketData.TicketsSold + 1

	// Get the tickets that corrolate with the eventID
	eKeyTickets, ok := k.GetTickets(ctx, eventID)
	if !ok {
		panic("No event with that ID")
	}
	// append the new ticket to the already existing ticket []
	eKeyTickets = append(eKeyTickets, ticket)
	k.SetTicket(ctx, k.eKey, eventID, eKeyTickets)

	// ownerKey toString to use as key in store
	ownerKey := ownerAddress.String()

	// append the newly created ticket to the users store
	var uTickets []emTypes.Ticket
	uTickets = append(uTickets, ticket)
	k.SetTicket(ctx, k.uKey, ownerKey, uTickets)
}

//Add the ticket to the market store
// func (k Keeper) ResaleTicket(ctx sdk.Context, ticketID string, eventID string) {
// 	ticket := k.GetTicket(ctx, eventID, ticketID)

// 	// get the [] of tickets in the market fo this event
// 	marketStore, ok := k.GetMarketPlaceTickets(ctx, eventID)
// 	if !ok {
// 		panic("no event in the marketplace")
// 	}

// 	// append the ticket to the existing [] of tickets
// 	marketStore = append(marketStore, ticket)
// 	k.SetTicket(ctx, k.mKey, eventID, marketStore)
// }

func (k Keeper) AddTicketToMarket(ctx sdk.Context, ticketID string, eventID string, markUp int) {
	ticket, ok := k.GetTicket(ctx, ticketID, eventID)
	if !ok {
		panic("Something")
	}
	requestedMarkUp := markUp
	var requested int64
	requested = int64(requestedMarkUp)

	ticket.SetNewPrice(requested)

	marketStore, ok := k.GetMarketPlaceTickets(ctx, eventID)
	if !ok {
		panic("no event in the marketplace")
	}
	marketStore = append(marketStore, ticket)

	k.SetTicket(ctx, k.mKey, eventID, marketStore)
}

func (k Keeper) SellTicket(ctx sdk.Context, ticketID string, eventID string,
	price int, newOwnerName string, newOwnerAddress sdk.AccAddress, ownerAddress sdk.AccAddress) {
	ticket := k.GetMarketPlaceTicket(ctx, eventID, ticketID)

	// check if the original owner is authorizing the transfer
	if !ownerAddress.Equals(ticket.OwnerAddress) {
		panic("Incorrect Owner")
	}

	ticket.ResaleTicket(newOwnerName, newOwnerAddress)

	eventTickets, ok := k.GetTickets(ctx, eventID)
	if !ok {
		panic("No ticket with that ID")
	}

	// find and replace the existing ticket
	for _, eTicket := range eventTickets {
		if eTicket.TicketID == ticket.TicketID {
			eTicket = ticket
		}
	}

	k.SetTicket(ctx, k.eKey, eventID, eventTickets)
}
