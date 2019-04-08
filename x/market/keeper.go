package market

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

//  Keeper for the market module
type Keeper struct {
	CKeeper bank.Keeper
	EKey    sdk.StoreKey // upcoming event key
	TKey    sdk.StoreKey // key for tickets that are generated for the people
	MKey    sdk.StoreKey // marketplace key for reselling
	UKey    sdk.StoreKey // store to keep an array of all the user tickets
	cdc     *codec.Codec
}

func NewKeeper(cKeeper bank.Keeper, eKey sdk.StoreKey, tKey sdk.StoreKey, mKey sdk.StoreKey, uKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CKeeper: cKeeper,
		EKey:    eKey,
		TKey:    tKey,
		MKey:    mKey,
		UKey:    uKey,
		cdc:     cdc,
	}
}

// func (k Keeper) GetTicket
// func (k Keeper) CreateTicket // add ticket to UKey store
// func (k Keeper) MoveTicketResale
// func (k Keeper) GetTickets within event
// func (k Keeper) GetNewPrice
// func (k Keeper) SellTicket // changeOwner
// func (k Keeper) ChangeOwner
