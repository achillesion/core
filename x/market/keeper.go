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
	eKey    sdk.StoreKey // upcoming event key
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

// Get all tickets of an event
func (k Keeper) GetTickets(ctx sdk.Context, eventID string) []emTypes.Ticket {
	store := ctx.KVStore(k.eKey)
	event := store.Get([]byte(eventID))
	var Tickets []emTypes.Ticket
	k.cdc.MustUnmarshalBinaryBare(event, &Tickets)
	return Tickets
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
}

// Get all tickets that a user may have
func (k Keeper) GetUserTickets(ctx sdk.Context, userAddress sdk.AccAddress) []emTypes.Ticket {
	store := ctx.KVStore(k.uKey)
	user := store.Get([]byte(userAddress))
	var Tickets []emTypes.Ticket
	k.cdc.MustUnmarshalBinaryBare(user, &Tickets)
	return Tickets
}

// initialPrice sdk.Coin, ticketNumber int, totalTickets int,
// markUpAllowed int, resale bool, price sdk.Coin
func (k Keeper) CreateTicket(ctx sdk.Context, parentReference string, ownerName string, ownerAddress sdk.AccAddress) emTypes.Ticket { // add ticket to UKey and EKey
	event := em.GetOpenEvent(ctx, parentReference)
	ticketData := event.TicketData
	ticket := emTypes.CreateTicket(ownerName, ownerAddress, parentReference,
		ticketData.InitialPrice, ticketData.TicketsSold ticketData.MarkUpAllowed,
		ticketData.Resale, ticketData.InitialPrice)
	ticketData.TicketNumber = ticketData.TicketNumber + 1
	return ticket
}

func (k Keeper) MoveTicketResale(ctx sdk.Context, ownerAddress sdk.AccAddress, ticketID string, eventID string) {

}

// func (k Keeper) SellTicket // changeOwner
