package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	ticketTypes "github.com/marbar3778/tic_mark/types"
	"github.com/spf13/cobra"
)

// TODO:

func GetCmdGetEvent(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get [event]",
		Short: "Get event",
		Long:  "Get a specific event",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			event := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/event/%s", queryRoute, event), nil)
			if err != nil {
				fmt.Println("could not resolve event name - %s \n", string(event))
				return nil
			}
			var eventData ticketTypes.Event
			cdc.MustUnmarshalJSON(res, &eventData)
			return cliCtx.PrintOutput(eventData)
		},
	}
}
