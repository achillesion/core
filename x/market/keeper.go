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
	cdc     *codec.Codec
}

func NewKeeper(cKeeper bank.Keeper, eKey sdk.StoreKey, tKey sdk.StoreKey, mKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CKeeper: cKeeper,
		EKey:    eKey,
		TKey:    tKey,
		MKey:    mKey,
		cdc:     cdc,
	}
}


func (k Keeper) GetTicket