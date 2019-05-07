package types

import (
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
