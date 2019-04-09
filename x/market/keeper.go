package market

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/marbar3778/tic_mark/x/eventmaker"
	emTypes "github.com/marbar3778/tic_mark/types"
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
	var Tickets []ticketType.Ticket
	k.cdc.MustUnmarshalBinaryBare(event, &Tickets)
	return Tickets
}

// Get all tickets that a user may have
func (k Keeper) GetUserTickets(ctx sdk.Context, userAddress sdk.AccAddress) []emTypes.Ticket {
	store := ctx.KVStore(k.uKey)
	user := store.Get([]byte(userAddress))
	vat Tickets []ticketType.Ticket
	k.cdc.MustUnmarshalBinaryBare(user, &Tickets)
	return Tickets
}

// Take the ticket data from the event and set it the ticket data
// ownerName string, ownerAddress sdk.AccAddress, parentReference string,
// 	initialPrice sdk.Coin, ticketNumber int, totalTickets int,
// 	markUpAllowed int, resale bool, price sdk.Coin
func (k Keeper) CreateTicket(ctx sdk.Context, parentReference string, ownerName string, ownerAddress sdk.AccAddress, ) emTypes.Ticket { // add ticket to UKey and EKey
	event := eventmaker.GetOpenEvent(ctx, parentReference)
	ticketData := event.TicketData
	emTypes.CreateTicket(ownerName, ownerAddress, parentReference, ticketData.InitialPrice, ticketData.TicketNumber)

}

// func (k Keeper) MoveTicketResale
// func (k Keeper) SellTicket // changeOwner
