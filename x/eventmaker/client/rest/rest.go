package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/marbar3778/tic_mark/x/eventmaker"

	"github.com/gorilla/mux"
)

const (
	restName = "event"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/event", storeName), createEventHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/event/{%s}", storeName, restName), getEventHandler(cdc, cliCtx, storeName)).Methods("GET")
}

type createEventReq struct {
	BaseReq           rest.BaseReq `json:"base_req"`
	EventName         string       `json:"event_name"`
	TotalTickets      int          `json:"total_tickets"`
	TicketsSold       int          `json:"tickets_sold"` // this most likely will be zero but if they are splitting tickets across platforms
	EventOwner        string       `json:"event_owner"`
	EventOwnerAddress string       `json:"event_owner_address"`
	Resale            bool         `json:"resale"`
}

// NewMsgCreateEvent(eventName string, totalTickets int, ticketsSold int, eventOwner string, eventOwnerAddress sdk.AccAddress, resale bool)
func createEventHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createEventReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.EventOwnerAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := eventmaker.NewMsgCreateEvent(req.EventName, req.TotalTickets, req.TicketsSold, req.EventOwner, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setNewOwnerReq struct {
	BaseReq              rest.BaseReq `json:"base_req"`
	EventName            string       `json:"event_name"`
	PreviousOwnerAddress string       `json:"previous_owner_address"`
	NewOwnerAddress      string       `json:"new_owner_address"`
	NewOwnerName         string       `json:"new_owner_name"`
}

//  NewMsgNewOwner(eventName string, previousOwnerAddress sdk.AccAddress, newOwnerAddress sdk.AccAddress, newOwner string)
func setNewOwnerHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setNewOwnerReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		prevAddr, err := sdk.AccAddressFromBech32(req.PreviousOwnerAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		newAddr, err := sdk.AccAddressFromBech32(req.NewOwnerAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := eventmaker.NewMsgNewOwner(req.EventName, prevAddr, newAddr, req.NewOwnerName)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

func getEventHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/event/%s", storeName, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

// TODO: get all events
