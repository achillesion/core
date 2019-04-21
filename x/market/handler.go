package market

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateTicket:
			return handleMsgCreateTicket(ctx, k, msg)
		case MsgAddTicketToMarket:
			return handleMsgAddTicketToMarket(ctx, k, msg)
		// case MsgSellTicket:
		// return handleMsgSellTicket(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized message: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCreateTicket(ctx sdk.Context, k Keeper, msg MsgCreateTicket) sdk.Result {
	k.CreateTicket(ctx, msg.EventID, msg.OwnerName, msg.OwnerAddress)
	return sdk.Result{}
}

func handleMsgAddTicketToMarket(ctx sdk.Context, k Keeper, msg MsgAddTicketToMarket) sdk.Result {
	k.ResaleTicket(ctx, msg.TicketID, msg.EventID)
	return sdk.Result{}
}
