package cli

import (
	"github.com/spf13/cobra"
	"net/http"
	"github.com/cosmos/cosmos-sdk/client/utils"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/marbar3778/tic_mark/x/market"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

type Ticket struct {
	EventID      string       `json:"event_id"`
	OwnerName    string       `json:"owner_name"`
	OwnerAddress string       `json:"owner_address"`
}

func GetCmdCreateTicket(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "createticket [eventID] [ownerName]",
		Short: "Create a Ticket",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}



			msg := market.NewMsgCreateTicket(args[0], args[1], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
}

func GetCmdResellTicket(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "ressellticket [eventID] [ticketID]",
		Short: "Try to resell your ticket",
		Args: cobra.ExactArgs(2),
		RunE: func (cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}
			msg := market.NewMsgResaleTicket(args[0], args[1], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			cliCtx.PrintResponse = true
			
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		}
	}
}