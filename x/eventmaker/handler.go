package eventmaker

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler : Handle messages to make changes to the store
func NewHandler(k BaseKeeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateEvent:
			return handleMsgCreateEvent(ctx, k, msg)
		case MsgNewOwner:
			return handleMsgNewOwner(ctx, k, msg)
		case MsgCloseEvent:
			return handleMsgCloseEvent(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized message: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCreateEvent(ctx sdk.Context, k BaseKeeper, msg MsgCreateEvent) sdk.Result {
	k.CreateEvent(ctx, msg.EventName, msg.TotalTickets, msg.EventOwner,
		msg.EventOwnerAddress, msg.Resale, msg.TicketData,
		msg.EventDetails)
	return sdk.Result{}
}

func handleMsgNewOwner(ctx sdk.Context, k BaseKeeper, msg MsgNewOwner) sdk.Result {
	if !msg.PreviousOwnerAddress.Equals(k.GetEventOwner(ctx, msg.EventName)) {
		return sdk.ErrUnauthorized(fmt.Sprintf("Unauthorized address: %s", msg.PreviousOwnerAddress)).Result()
	}
	k.NewOwner(ctx, msg.EventName, msg.PreviousOwnerAddress, msg.NewOwnerAddress, msg.NewOwner)
	return sdk.Result{}
}

func handleMsgCloseEvent(ctx sdk.Context, k BaseKeeper, msg MsgCloseEvent) sdk.Result {
	if !msg.EventOwnerAddress.Equals(k.GetEventOwner(ctx, msg.EventID)) {
		return sdk.ErrUnauthorized(fmt.Sprintf("Unauthorized address: %s", msg.EventOwnerAddress)).Result()
	}
	k.CloseEvent(ctx, msg.EventID)
	return sdk.Result{}
}
