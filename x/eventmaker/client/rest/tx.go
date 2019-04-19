package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/marbar3778/tic_mark/types"
	"github.com/marbar3778/tic_mark/x/eventmaker"
)

type createEventReq struct {
	BaseReq           rest.BaseReq       `json:"base_req"`
	EventName         string             `json:"event_name"`
	TotalTickets      int                `json:"ticket_owner"`
	EventOwner        string             `json:"event_owner"`
	EventOwnerAddress string             `json:"event_owner_address"`
	Resale            bool               `json:"resale"`
	TicketData        types.TicketData   `json:"ticket_data"`
	EventDetails      types.EventDetails `json:"event_details"`
}

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

		msg := eventmaker.NewMsgCreateEvent(req.EventName, req.TotalTickets, req.EventOwner, addr, req.Resale, req.TicketData, req.EventDetails)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type closeEventReq struct {
	BaseReq           rest.BaseReq `json:"base_req"`
	EventID           string       `json:"event_id"`
	EventOwnerAddress string       `json:"owner_address"`
}

func closeEvent(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req closeEventReq
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
		msg := eventmaker.NewMsgCloseEvent(req.EventID, addr)
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
