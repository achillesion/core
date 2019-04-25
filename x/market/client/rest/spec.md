## Txs

```golang
type addTicketToMarketReq struct {
	BaseReq            rest.BaseReq `json:"base_req"`
	EventID            string       `json:"event_id"`
	TicketID           string       `json:"ticket_id"`
	TicketOwnerAddress string       `json:"ticket_owner_address"`
	SalePrice          int          `json:"sale_price"`
}
```

```golang
type createTicketReq struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	EventID      string       `json:"event_id"`
	OwnerName    string       `json:"owner_name`
	OwnerAddress string       `json:"owner_address"`
}
```

## Query

`/%s/tickets/{EventID}`
`/%s/%s/ticket/{TicketID}`
