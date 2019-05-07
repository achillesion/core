package types

import (
	"fmt"
	"os/exec"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Ticket struct {
	TicketID        string         // UUID for the ticket
	OwnerName       string         // owner of the item
	OwnerAddress    sdk.AccAddress // owner address
	ParentReference string         // reference to parent in this case a event, UUID of the parent
	InitialPrice    sdk.Coin       // original price of the item, if initialPrice is 0 then its a free event
	TicketNumber    int            // if the parent wants to make more than one
	TotalTickets    int            // to give the user a sense of power that they are the only one with this number
	MarkUpAllowed   int64          // amount of the current price (originalPrice || newPrice)
	Resale          bool           // if the ticket is allowed to enter the market place
	ResaleCounter   int            // amount of times it the item has been resold
	Price           sdk.Coin       // price that the item will be resold for
}

func CreateTicket(ownerName string, ownerAddress sdk.AccAddress, parentReference string,
	initialPrice sdk.Coin, ticketNumber int, totalTickets int,
	markUpAllowed int64, resale bool, price sdk.Coin) Ticket {

	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err)
	}

	uuid := fmt.Sprintf("%s", out)

	return Ticket{
		TicketID:        uuid,
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

// 1 tic = 100htic

// Set new price
func (t Ticket) getNewPrice(markUp int64) sdk.Coin { // TODO: Fixme
	percent := int64(100)
	Amount := t.Price.Amount.Int64()

	markUpAmount := ((markUp * percent) / Amount)

	fmt.Println(markUpAmount)

	newAmount := sdk.NewInt(markUpAmount)

	t.Price.Amount = t.Price.Amount.Add(newAmount)

	return t.Price
}

func (t Ticket) ResaleTicket(ownerName string, ownerAddress sdk.AccAddress) {
	t.OwnerName = ownerName
	t.OwnerAddress = ownerAddress

	// t.Price = t.SetNewPrice(markUp)
}

// Get the new price of the ticket for resale
func (t Ticket) SetNewPrice(markUp int64) Ticket {
	if !t.Resale {
		panic("Can not enter the marketplace")
	}
	if markUp > t.MarkUpAllowed {
		panic("The markup suggested is to great")
	}

	t.ResaleCounter++
	t.Price = t.getNewPrice(markUp)

	return t
}

// Set new owner TODO: make changes be immutable, spawn a new ticket
func (t Ticket) SetNewOwner(ownerName string, ownerAddress sdk.AccAddress) Ticket {
	t.OwnerName = ownerName
	t.OwnerAddress = ownerAddress
	return t
}

// Get my ticket number
func (t Ticket) GetTicketNumber() string {
	return fmt.Sprintf("Ticket: %d/%d", t.TicketNumber, t.TotalTickets)
}

func (t Ticket) String() string {
	return fmt.Sprintf("TicketID: %s", t.TicketID)
}
