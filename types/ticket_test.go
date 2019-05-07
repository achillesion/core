package types

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

var coin = sdk.Coin{Denom: "tic", Amount: amount}

var ticket = Ticket{
	"123",
	"Bob",
	sdk.AccAddress{byte('A')},
	"parent",
	coin,
	5,
	5,
	percent,
	true,
	0,
	coin,
}

func TestTicketCreation(t *testing.T) {
	createdTicket := CreateTicket("Bob", sdk.AccAddress{byte('A')}, "parent", coin, 5, 5, percent, true, coin)

	assert.Equal(t, createdTicket.OwnerName, ticket.OwnerName, "Ownernames should be equal")
	assert.Equal(t, createdTicket.OwnerAddress, ticket.OwnerAddress, "Owner addresses should be equal")
	assert.Equal(t, createdTicket.ParentReference, ticket.ParentReference, "Parnet References should be equal")
	assert.Equal(t, createdTicket.InitialPrice, ticket.InitialPrice, "Prices should be equal")
}

func TestGetTicketNumber(t *testing.T) {
	createdTicket := CreateTicket("Bob", sdk.AccAddress{byte('A')}, "parent", coin, 5, 5, percent, true, coin)

	ticketNumber := createdTicket.GetTicketNumber()

	assert.Equal(t, ticketNumber, fmt.Sprintf("Ticket: %d/%d", ticket.TicketNumber, ticket.TotalTickets))
}

func TestSetNewOwner(t *testing.T) {
	createdTicket := CreateTicket("Bob", sdk.AccAddress{byte('A')}, "parent", coin, 5, 5, percent, true, coin)

	newAddress := sdk.AccAddress{byte('B')}

	newTicket := createdTicket.SetNewOwner("Marko", newAddress)

	assert.Equal(t, newTicket.OwnerAddress, newAddress, "Addresses should be equal")
	assert.Equal(t, newTicket.OwnerName, "Marko", "Addresses should be equal")
}

func TestSetNewPrice(t *testing.T) {
	createdTicket := CreateTicket("Bob", sdk.AccAddress{byte('A')}, "parent", coin, 5, 5, percent, true, coin)

	eight := int64(8)
	sixteen := int64(16)

	proposedAmountEight := sdk.Coin{Denom: "tic", Amount: sdk.NewInt(int64(108))}
	proposedAmount16 := sdk.Coin{Denom: "tic", Amount: sdk.NewInt(int64(116))}

	newAmountEight := createdTicket.SetNewPrice(eight)
	assert.Equal(t, newAmountEight.Price, proposedAmountEight, "Should be 108tic")
	assert.Equal(t, newAmountEight.ResaleCounter, 1, "Resale Counter should be 1")

	newAmount16 := createdTicket.SetNewPrice(sixteen)
	assert.Equal(t, newAmount16.Price, proposedAmount16, "Should be 116tic")
	assert.Equal(t, newAmount16.ResaleCounter, 1, "Resale Counter should be 1")

	newAmount24 := newAmount16.SetNewPrice(eight)
	fmt.Println(newAmount24.Price) // TODO: wrong float conversion
}
