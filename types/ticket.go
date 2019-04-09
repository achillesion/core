package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Ticket struct {
	OwnerName       string         // owner of the item
	OwnerAddress    sdk.AccAddress // owner address
	ParentReference string         // reference to parent in this case a event, UUID of the parent
	InitialPrice    sdk.Coin       // original price of the item, if initialPrice is 0 then its a free event
	TicketNumber    int            // if the parent wants to make more than one
	TotalTickets    int            // to give the user a sense of power that they are the only one with this number
	MarkUpAllowed   int            // amount of the current price (originalPrice || newPrice)
	Resale          bool           // if the ticket is allowed to enter the market place
	ResaleCounter   int            // amount of times it the item has been resold
	Price           sdk.Coin       // price that the item will be resold for
}

func CreateTicket(ownerName string, ownerAddress sdk.AccAddress, parentReference string,
	initialPrice sdk.Coin, ticketNumber int, totalTickets int,
	markUpAllowed int, resale bool, price sdk.Coin) Ticket {
	return Ticket{
		OwnerName:       ownerName,
		OwnerAddress:    ownerAddress,
		ParentReference: parentReference,
		InitialPrice:    initialPrice,
		TicketNumber:    ticketNumber,
		TotalTickets:    totalTickets,
		MarkUpAllowed:   markUpAllowed,
		Resale:          resale,
		ResaleCounter:   0,
		Price:           price,
	}
}

// Set new price
func (t Ticket) SetNewPrice(oldPrice sdk.Coin, markUp int) sdk.Coin {

	// maxMarkUp := markUp / 100
	// markUpAmount := oldPrice.Amount * maxMarkUp
	return oldPrice
}

// Get the new price of the ticket for resale
func (t Ticket) GetMaxNewPrice(markUp int) sdk.Coin {
	if !t.Resale {
		panic("Can not enter the marketplace")
	}
	if t.ResaleCounter > 1 {
		t.Price.Add(t.SetNewPrice(t.InitialPrice, markUp))
		t.ResaleCounter++
		return t.Price
	}
	t.ResaleCounter++
	t.Price.Add(t.SetNewPrice(t.Price, markUp))
	return t.Price
}

// Get the current owner of the ticker
func (t Ticket) GetCurrentOwner() string {
	return fmt.Sprintf("Ticket Owner: %s, Ticket Owner Address: %s", t.OwnerName, t.OwnerAddress.String())
}

// Set new owner
func (t Ticket) SetNewOwner(ownerName string, ownerAddress sdk.AccAddress) string {
	t.OwnerName = ownerName
	t.OwnerAddress = ownerAddress
	return fmt.Sprintf("New Ticket Owner: %s, New Ticket Owner Address: %s", t.OwnerName, t.OwnerAddress)
}

// Get my ticket number
func (t Ticket) GetTicketNumber() string {
	return fmt.Sprintf("Ticket: %d/%d", t.TicketNumber, t.TotalTickets)
}
