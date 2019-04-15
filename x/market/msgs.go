package market

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = "market"

type MsgCreateTicket struct {
	EventID      string         `json:"event_id"`
	OwnerName    string         `json:"owner_name`
	OwnerAddress sdk.AccAddress `json:"owner_address"`
}

func NewMsgCreateTicket(eventID string, ownerName string, ownerAddress sdk.AccAddress) MsgCreateTicket {
	return MsgCreateTicket{
		EventID:      eventID,
		OwnerName:    ownerName,
		OwnerAddress: ownerAddress,
	}
}

//nolint
func (msg MsgCreateTicket) Route() string { return RouterKey }
func (msg MsgCreateTicket) Type() string  { return "create_ticket" }
func (msg MsgCreateTicket) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.OwnerAddress}
}

func (msg MsgCreateTicket) ValidateBasic() sdk.Error {
	if len(msg.EventID) == 0 {
		return sdk.ErrUnknownRequest("There is no eventID")
	}
	if len(msg.OwnerName) == 0 {
		return sdk.ErrUnknownRequest("The owner name is not present")
	}
	if msg.OwnerAddress.Empty() {
		return sdk.ErrInvalidAddress("Missing the owners address")
	}
	return nil
}

func (msg MsgCreateTicket) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

type MsgResaleTicket struct {
	EventID     string         `json:"event_id"`
	TicketID    string         `json:"ticket_id"`
	TicketOwner sdk.AccAddress `json:"ticket_owner"`
}

func NewMsgResaleTicket(eventID string, ticketID string, ticketOwner sdk.AccAddress) MsgResaleTicket {
	return MsgResaleTicket{
		EventID:     eventID,
		TicketID:    ticketID,
		TicketOwner: ticketOwner,
	}
}

//nolint
func (msg MsgResaleTicket) Route() string { return RouterKey }
func (msg MsgResaleTicket) Type() string  { return "resale_ticket" }
func (msg MsgResaleTicket) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.TicketOwner}
}

func (msg MsgResaleTicket) ValidateBasic() sdk.Error {
	if len(msg.EventID) == 0 || len(msg.TicketID) == 0 {
		return sdk.ErrUnknownRequest(fmt.Sprintf("There is no eventID and/or ticketID, eventID: %s, ticketID: %s", msg.EventID, msg.TicketID))
	}
	if msg.TicketOwner.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf("Please provide a valid address, current address: %s", msg.TicketOwner))
	}
	return nil
}

func (msg MsgResaleTicket) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}
