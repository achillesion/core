# Event Generator Module

This module is set to create the Events that are needed.

## Create Event Data Needed

```golang
type Event struct {
  EventName         string
  EventOwner        string
  EventOwnerAddress sdk.AccAddress
  TicketData        TicketData
  EventDetails      EventDetails
}
```

```golang
type EventDetails struct {
  Address      string
  City         string
  Country      string
  Date         string // openingDate - closingDate
}
```

```golang
type TicketData struct {
  Resale        bool
  TotalTickets  int
  TicketsSold   int
  InitialPrice  sdk.Coins
  Resale        bool
  MarkUpAllowed int
}
```

### Uses

- Create Paid & Free Event
- Get Event
- Get list of Events
- Event has escrow to hold money
- Close event

### User Flows

- As an event owner I want to create a free event
- As an event owner I want to create a paid event
- As an event owner I want to change ownership of my event to another owner
- As an event owner I want to give someone else, of my choosing, access to the data produced from sales
- As an event owner I want to close my event, because it has passed
- As an event owner I want to close my event because it was cancelled
- As a user I want a ticket from a specific event
