package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	ticketTypes "github.com/marbar3778/tic_mark/x/eventmaker/types"
	"github.com/spf13/cobra"
)

func GetCmdGetTickets(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get tickets",
		Short: "tickets",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/tickets", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query tickets\n")
				return nil
			}

			var tickets ticketTypes.Ticket // has to be an array not a single ticket
			cdc.MustUnmarshalJSON(res, &tickets)
			return cliCtx.PrintOutput(tickets)
		},
	}
}

func GetCmdGetTicket(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "Get ticket [ticketID]",
		Short: "Get specific ticket",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			event := args[0]
			ticketID := args[1]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/event/%s/ticket/%s", queryRoute, event, ticketID), nil)
			if err != nil {
				fmt.Println("could not resolve event name: %s, ticket: %s \n", event, ticketID)
				return nil
			}
			var ticket ticketTypes.Ticket
			cdc.MustUnmarshalJSON(res, &ticket)
			return cliCtx.PrintOutput(ticket)
		},
	}
}
