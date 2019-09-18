package types

import (
	"fmt"
	"os/exec"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TicketData struct {
	InitialPrice  sdk.Coin `json:"ticket_price"`    // Price of this ticket
	MarkUpAllowed int      `json:"mark_up_allowed"` // if ticket can be sold, what is max amount over originalPrice user can set per sale
	TotalTickets  int      `json:"total_tickets"`   // Total amount of tickets
	TicketsSold   int      `json:"tickets_sold"`    // amount of tickets sold
	Resale        bool     `json:"resale"`          // Mark if secondary market is possible
}

type EventDetails struct {
	LocationName string `json:"location_name"`
	Address      string `json:"address"` // Address of event
	City         string `json:"city"`    // City in which the event is
	Country      string `json:"country"` // Country the event is being held in
	Date         string `json:"date"`    // date of the event, TODO: make it enable a multi day event
}

type Event struct {
	EventID           string         `json:"event_id"`            // UUID for event
	EventName         string         `json:"event_name"`          // Name of the Event
	EventOwner        string         `json:"event_owner"`         // Event Organizer
	EventOwnerAddress sdk.AccAddress `json:"event_owner_address"` // Event Organizer Address
	TicketData        TicketData     `json:"ticket_data"`         // From which data to generate tickets
	EventDetail       EventDetails   `json:"event_details"`       // struct containing the event details
}

// CreateEvent, creates event of the organizers
func CreateEvent(eventName string, totalTickets int, eventOwner string,
	eventOwnerAddress sdk.AccAddress, resale bool, ticketData TicketData,
	eventDetails EventDetails) Event {

	if totalTickets <= 0 {
		panic(fmt.Sprintf("amount of tickets can not be zero or less(-), you're amount is %v", totalTickets))
	}
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err)
	}

	uuid := fmt.Sprintf("%s", out)

	return Event{
		EventID:           uuid,
		EventName:         eventName,
		EventOwner:        eventOwner,
		EventOwnerAddress: eventOwnerAddress,
		TicketData:        ticketData,
		EventDetail:       eventDetails,
	}
}

// Check if eventName, eventOwner & eventOwnerAddress are not null
func (e Event) ValidEventCreation() {
	if e.EventName == "" || e.EventOwner == "" || e.EventOwnerAddress == nil {
		panic("Event does not have a name")
	}
}

func (e Event) String() string {
	return fmt.Sprintf(
		`EventName: %s
		EventOwner: %s
		EventOwnerAddress: %s
		TicketData: %v
		EventDetail: %v`, e.EventName, e.EventOwner, e.EventOwnerAddress.String(),
		e.TicketData, e.EventDetail)
}

// Eventdetails check
func (e Event) EventDetails() string {
	return fmt.Sprintf("Event name:%s, Total amount of tickets:%d, Event owner: %s", e.EventName, e.TicketData.TotalTickets, e.EventOwner)
}

// GetTicketDetails - Get details of ticket to create tickets with
func (e Event) GetTicketDetails() TicketData {
	return e.TicketData
}

// GetEventDetails - get event details
func (e Event) GetEventDetails() EventDetails {
	return e.EventDetail
}

func (e Event) SetDate(date string) EventDetails {
	e.EventDetail.Date = date
	return e.EventDetail
}
