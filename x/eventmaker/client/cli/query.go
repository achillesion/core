package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	ticketTypes "github.com/marbar3778/tic_mark/x/eventmaker/types"
	em "github.com/marbar3778/tic_mark/x/eventmaker"
	"github.com/spf13/cobra"
)

unc (mc ModuleClient) GetQueryCmd() *cobra.Command {
	ticketQueryCmd := &cobra.Command{
		Use:   "eventmaker",
		Short: "Querying commands for the eventmaker module",
	}
	ticketQueryCmd.AddCommand(client.GetCommands(
		emd.GetCmdGetOpenEvent(mc.storekey, mc.cdc),
		emd.GetCmdGetClosedEvent(mc.storekey, mc.cdc),
		emd.GetCmdGetOwner(mc.storekey, mc.cdc),
	)...)

	return ticketQueryCmd
}

func GetCmdGetOpenEvent(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "openevent [Event]",
		Short: "Get an Open Event",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			event := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/openevent/%s", queryRoute, event), nil)
			if err != nil {
				fmt.Printf("could not resolve event name - %s \n", event)
				return nil
			}
			var eventData ticketTypes.Event
			cdc.MustUnmarshalJSON(res, &eventData)
			return cliCtx.PrintOutput(eventData)
		},
	}
}

func GetCmdGetClosedEvent(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "closedevent [Event]",
		Short: "Get an Closed Event",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			event := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/closedevent/%s", queryRoute, event), nil)
			if err != nil {
				fmt.Printf("could not resolve event name - %s \n", event)
				return nil
			}
			var eventData ticketTypes.Event
			cdc.MustUnmarshalJSON(res, &eventData)
			return cliCtx.PrintOutput(eventData)
		},
	}
}

func GetCmdGetOwner(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "eventowner [Event]",
		Short: "Get the owner of the event",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			event := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/owner/%s", queryRoute, event), nil)
			if err != nil {
				fmt.Printf("could not resolve event name - %s \n", event)
				return nil
			}
			var owner em.QueryResOwner
			cdc.MustUnmarshalJSON(res, &owner)
			return cliCtx.PrintOutput(owner)
		},
	}
}
