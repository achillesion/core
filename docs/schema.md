# Schema for metadata

```
user {
  userID: string
  email: string
  hashedPW string
}
```

```
Tickets {
  TicketID: string
  userID: string
  eventID: string
  TicketCategory: string
  events: Events
}
```

```
ticketData {
  eventID: string
  ticketPrice: uint
  markUpAllowed: uint
  totalTickets: uint
  ticketsSold: uint
  resale: bool
}
```

```
Events {
  eventID: string
  eventName: string
  eventOwner: string
  eventOwnerAddress: string
  eventDetails: EventDetails
  ticketData: tickets
}
```

```
EventDetails {
  eventID: string
  locationName: string
  address: string
  city: string
  country: string
  date: string
}
```