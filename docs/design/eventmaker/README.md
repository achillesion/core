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
- Close event and refund everyone if event is cancelled
- Close event and withdraw money after event is successfully over

### User Flows

- As an event owner I want to create an event specifying:
  - the name of the event
  - date & time of the event
  - location & address of the venue
  - if its a free or a paid event
  - price in case its a paid event
  - enable or disable resale
  - price cap in case of resale
  - set percentage of comission
- As an event owner I want to change ownership of my event to another owner
- As an event owner I want to access and evaluate the data from the ticket life cycle
- As an event owner I want to give someone else, of my choosing, access to the data produced from sales
- As an event owner I want to close my event, because it has passed and withdraw the money from the escrow
- As an event owner I want to close my event because it was cancelled and refund all the ticket buyers
- As a user I want a ticket from a specific event
