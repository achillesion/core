package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/gorilla/mux"
)

const (
	restName = "market"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/tickets/{%s}", storeName, restName), getTicketsEventHandler(cdc, cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/%s/ticket/{TicketID}", storeName, restName), getTicketEventHandler(cdc, cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/createTicket", storeName), createTicketEventHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/addtickettomarket", storeName), addTicketToMarket(cdc, cliCtx)).Methods("POST")
}
