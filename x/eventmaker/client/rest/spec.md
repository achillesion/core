### Txs

```golang
 type createEventReq struct {
BaseReq rest.BaseReq `json:"base_req"`
EventName string `json:"event_name"`
TotalTickets int `json:"ticket_owner"`
EventOwner string `json:"event_owner"`
EventOwnerAddress string `json:"event_owner_address"`
Resale bool `json:"resale"`
TicketData types.TicketData `json:"ticket_data"`
EventDetails types.EventDetails `json:"event_details"`
}
```

```golang
 type closeEventReq struct {
	BaseReq           rest.BaseReq `json:"base_req"`
	EventID           string       `json:"event_id"`
	EventOwnerAddress string       `json:"owner_address"`
}
```

```golang
type setNewOwnerReq struct {
	BaseReq              rest.BaseReq `json:"base_req"`
	EventName            string       `json:"event_name"`
	PreviousOwnerAddress string       `json:"previous_owner_address"`
	NewOwnerAddress      string       `json:"new_owner_address"`
	NewOwnerName         string       `json:"new_owner_name"`
}
```

## query

`/%s/event/open/{%s}` second %s is eventID
`/%s/event/closed/{%s}` second %s is eventID
