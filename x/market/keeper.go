package market

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	emTypes "github.com/marbar3778/tic_mark/types"
	em "github.com/marbar3778/tic_mark/x/eventmaker"
)

//  Keeper for the market module
type Keeper struct {
	cKeeper bank.Keeper
	eKey    sdk.StoreKey // upcoming event key where the tickets will be held
	mKey    sdk.StoreKey // marketplace key for reselling
	uKey    sdk.StoreKey // store to keep an array of all the users that have tickets
	cdc     *codec.Codec
}

func NewKeeper(cKeeper bank.Keeper, eKey sdk.StoreKey, mKey sdk.StoreKey, uKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		cKeeper: cKeeper,
		eKey:    eKey,
		mKey:    mKey,
		uKey:    uKey,
		cdc:     cdc,
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
		panic("no ticket")
	}
	return emTypes.Ticket{}
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
	event := em.GetOpenEvent(ctx, eventID)
	ticketData := event.TicketData
	ticket := emTypes.CreateTicket(ownerName, ownerAddress, eventID,
		ticketData.InitialPrice, ticketData.TicketsSold, ticketData.TotalTickets, ticketData.MarkUpAllowed,
		ticketData.Resale, ticketData.InitialPrice)
	ticketData.TicketNumber = ticketData.TicketNumber + 1
	k.SetTicket(ctx, k.eKey, eventID, ticket) // set ticket to event store
	k.SetTicket(ctx, k.uKey, eventID, ticket) // set ticket to event store
}

// Add the ticket to the market store
func (k Keeper) ResaleTicket(ctx sdk.Context, ticketID string, eventID string) {
	ticket := k.GetTicket(ctx, eventID, ticketID)
	k.SetTicket(ctx, k.mKey, eventID, ticket)
}

func (k Keeper) SellTicket(ctx sdk.Context, ticketID string, eventID string,
	newOwnerName string, newOwnerAddress sdk.AccAddress, sellingPrice int) {
	ticket := k.GetTicket(ctx, ticketID, eventID)
	ticket.ResaleTicket(newOwnerName, newOwnerAddress, sellingPrice)
	k.SetTicket(ctx, k.eKey, eventID, ticket)
	// uStore := ctx.KVStore(k.uKey)
	// uStore.Delete make it delete a single entry of the key not the key
	mStore := ctx.KVStore(k.mKey)
	mStore.Delete([]byte(ticketID))
} // changeOwner
