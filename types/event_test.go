package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/go-cmp/cmp"
	"testing"
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
	eventData := CreateEvent("ME", 10, "You", sdk.AccAddress{byte('A')}, true, dataTicket, detailsEvent)
	if event.EventName != eventData.EventName {
		t.Errorf("Error Expected %s got %s", event.EventName, eventData.EventName)
	}
	if event.EventOwner != eventData.EventOwner {
		t.Errorf("Error Expected %s got %s", event.EventOwner, eventData.EventOwner)
	}
	if !cmp.Equal(event.EventOwnerAddress, eventData.EventOwnerAddress) {
		t.Errorf("Error Expected %s got %s", event.EventOwnerAddress, eventData.EventOwnerAddress)
	}
	if event.TicketData != eventData.TicketData {
		t.Error("Error Expected")
	}
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
	if ticketDets != dataTicket {
		t.Error("Expected the above ticket data")
	}
}

func TestEventDetails(t *testing.T) {
	eventData := CreateEvent("ME", 10, "You", sdk.AccAddress{byte('A')}, true, dataTicket, detailsEvent)
	eventDets := eventData.GetEventDetails()
	if eventDets != detailsEvent {
		t.Error("Expected the above event data")
	}
}

func TestSetDate(t *testing.T) {
	eventData := CreateEvent("ME", 10, "You", sdk.AccAddress{byte('A')}, true, dataTicket, detailsEvent)
	newDate := "01.01.2021"
	date := eventData.SetDate(newDate)
	if date.Date != newDate {
		t.Errorf("Expected: %s, Got: %s", newDate, eventData.EventDetail.Date)
	}
}
