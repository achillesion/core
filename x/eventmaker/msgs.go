package eventmaker

import (
	"encoding/json"
	"fmt"

	"github.com/marbar3778/tic_mark/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = "eventmaker"

// MsgCreateEvent : Msg to create event
type MsgCreateEvent struct {
	EventName         string             `json:"event_name"`
	TotalTickets      int                `json:"ticket_owner"`
	EventOwner        string             `json:"event_owner"`
	EventOwnerAddress sdk.AccAddress     `json:"event_owner_address"`
	Resale            bool               `json:"resale"`
	TicketData        types.TicketData   `json:"ticket_data"`
	EventDetails      types.EventDetails `json:"event_details"`
}

// NewMsgCreateEvent : create Event
func NewMsgCreateEvent(
	eventName string, totalTickets int,
	eventOwner string, eventOwnerAddress sdk.AccAddress, resale bool,
	ticketData types.TicketData, eventDetails types.EventDetails) MsgCreateEvent {
	return MsgCreateEvent{
		EventName:         eventName,
		TotalTickets:      totalTickets,
		EventOwner:        eventOwner,
		EventOwnerAddress: eventOwnerAddress,
		Resale:            resale,
		TicketData:        ticketData,
		EventDetails:      eventDetails,
	}
}

//nolint
func (msg MsgCreateEvent) Route() string { return RouterKey }
func (msg MsgCreateEvent) Type() string  { return "create_event" }
func (msg MsgCreateEvent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.EventOwnerAddress}
}

// ValidateBasic : validity check
func (msg MsgCreateEvent) ValidateBasic() sdk.Error {
	if len(msg.EventName) == 0 {
		return sdk.ErrUnknownRequest("There is no name to the event")
	}
	if len(msg.EventOwner) == 0 || msg.EventOwnerAddress.Empty() {
		return sdk.ErrUnknownRequest("The event does not have a owner")
	}
	if msg.TotalTickets == 0 {
		return sdk.ErrUnknownRequest("There are no tickets to be sold")
	}
	return nil
}

// GetSignBytes - receive message bytes for signer to sign
func (msg MsgCreateEvent) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// MsgNewOwner : set new owner of the event constructor
type MsgNewOwner struct {
	EventName            string         `json:"event_name"`
	PreviousOwnerAddress sdk.AccAddress `json:"previous_owner_address"`
	NewOwnerAddress      sdk.AccAddress `json:"new_owner_address"`
	NewOwner             string         `json:"new_owner"`
}

// NewMsgNewOwner : Place a new owner onto the event
func NewMsgNewOwner(eventName string, previousOwnerAddress sdk.AccAddress, newOwnerAddress sdk.AccAddress, newOwner string) MsgNewOwner {
	return MsgNewOwner{
		EventName:            eventName,
		PreviousOwnerAddress: previousOwnerAddress,
		NewOwnerAddress:      newOwnerAddress,
		NewOwner:             newOwner,
	}
}

//nolint
func (msg MsgNewOwner) Route() string { return RouterKey }
func (msg MsgNewOwner) Type() string  { return "new_owner" }
func (msg MsgNewOwner) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.PreviousOwnerAddress}
}

// ValidateBasic : validity check
func (msg MsgNewOwner) ValidateBasic() sdk.Error {
	if len(msg.EventName) == 0 {
		return sdk.ErrUnknownRequest("There is no event name")
	}
	if msg.PreviousOwnerAddress.Empty() || msg.NewOwnerAddress.Empty() {
		return sdk.ErrInvalidAddress("Missing previous owner addres or new owner address")
	}
	if len(msg.NewOwner) == 0 {
		return sdk.ErrUnknownRequest("Please provide a name for the new Owner")
	}
	return nil
}

// GetSignBytes - receive message bytes for signer to sign
func (msg MsgNewOwner) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// MsgCloseEvent : struct to close an event
type MsgCloseEvent struct {
	EventID           string         `json:"eventid"`
	EventOwnerAddress sdk.AccAddress `json:"event_owner_address"`
}

func NewMsgCloseEvent(eventID string, eventOwnerAddress sdk.AccAddress) MsgCloseEvent {
	return MsgCloseEvent{
		EventID:           eventID,
		EventOwnerAddress: eventOwnerAddress,
	}
}

//nolint
func (msg MsgCloseEvent) Route() string { return RouterKey }
func (msg MsgCloseEvent) Type() string  { return "close_event" }
func (msg MsgCloseEvent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.EventOwnerAddress}
}

// ValidateBasic : validity check
func (msg MsgCloseEvent) ValidateBasic() sdk.Error {
	if len(msg.EventID) == 0 {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Please Provide a valid event name: %s", msg.EventID))
	}
	if msg.EventOwnerAddress.Empty() {
		return sdk.ErrInvalidAddress("missing event owner address")
	}
	return nil
}

func (msg MsgCloseEvent) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}
