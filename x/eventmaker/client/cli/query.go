package cli

import (
	"fmt"

	ticketTypes "github.com/marbar3778/tic_mark/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetCmdGetEvent(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get [event]",
		Short: "Get event",
		Long:  "Get a specific event",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewClIContext().WithCodec(cdc)
			event := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/event/%s", queryRoute, event), nil)
			if err != nil {
				fmt.Printf("could not resolve event name - %s \n", string(event))
				return nil
			}
			var data ticketTypes.Event
			cdc.MustUnmarshalJSON(res, &data)
			return cliCtx.PrintOutPut(data)
		},
	}
}
