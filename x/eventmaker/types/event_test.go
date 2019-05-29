package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

var oneHundred = 100
var percent = int64(oneHundred)
var amount = sdk.NewInt(percent)

// create ticketData
var dataTicket = TicketData{
	sdk.Coin{Denom: "tic", Amount: amount},
	5,
	10,
	0,
	true,
}

// create eventData
var detailsEvent = EventDetails{
	"Da House",
	"4239 sw 4th ave",
	"Portland, OR",
	"USA",
	"01/02/2020",
}

// create Event
var event = Event{
	"1234",
	"ME",
	"You",
	sdk.AccAddress{byte('A')},
	dataTicket,
	detailsEvent,
}

func TestCreateEvent(t *testing.T) {
	eventData := CreateEvent("ME", 10, "You", v, true, dataTicket, detailsEvent)

	assert.Equal(t, event.EventName, eventData.EventName, "they should be equal")

	assert.Equal(t, event.EventOwner, eventData.EventOwner, "They should be equal")

	assert.Equal(t, event.EventOwnerAddress, eventData.EventOwnerAddress, "Equal Addresses")

	assert.Equal(t, event.TicketData, eventData.TicketData, "Ticket Data should be equal")
}

func TestValidEventCreation(t *testing.T) {
	eventData := CreateEvent("ME", 10, "You", sdk.AccAddress{byte('A')}, true, dataTicket, detailsEvent)
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code did not panic")
		}
	}()
	eventData.ValidEventCreation()
}

func TestTicketDetails(t *testing.T) {
	eventData := CreateEvent("ME", 10, "You", sdk.AccAddress{byte('A')}, true, dataTicket, detailsEvent)
	ticketDets := eventData.GetTicketDetails()

	assert.Equal(t, ticketDets, dataTicket, "Expected the ticket data to be equal")
}

func TestEventDetails(t *testing.T) {
	eventData := CreateEvent("ME", 10, "You", sdk.AccAddress{byte('A')}, true, dataTicket, detailsEvent)
	eventDets := eventData.GetEventDetails()

	assert.Equal(t, eventDets, detailsEvent, "event details should be equal")
}

func TestSetDate(t *testing.T) {
	eventData := CreateEvent("ME", 10, "You", sdk.AccAddress{byte('A')}, true, dataTicket, detailsEvent)
	newDate := "01.01.2021"
	date := eventData.SetDate(newDate)

	assert.Equal(t, date.Date, newDate, "Dates should be equal")
}
