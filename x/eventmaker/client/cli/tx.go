package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	types "github.com/marbar3778/tic_mark/x/eventmaker/types"
	em "github.com/marbar3778/tic_mark/x/eventmaker"
)

func (mc ModuleClient) GetTxCmd() *cobra.Command {
	ticketTxCmd := &cobra.Command{
		Use:   "eventmaker",
		Short: "eventmaker tx subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	ticketTxCmd.AddCommand(client.PostCommands(
		emd.GetCmdCreateEvent(mc.cdc),
		emd.GetCmdNewOwner(mc.cdc),
		emd.GetCmdCloseEvent(mc.cdc),
	)...)

	return ticketTxCmd
}

// CreateEvent(ctx sdk.Context, eventName string, totalTickets int,
// 	eventOwner string, eventOwnerAddress sdk.AccAddress, resale bool,
// 	ticketData ticType.TicketData, eventDetails ticType.EventDetails)
func GetCmdCreateEvent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createEvent [eventName] [totalTickets] [eventOwner] [resale] [ticketData] [eventDetails]",
		Short: "Create Event",
		Long:  "Create a event",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			num, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			boo, err := strconv.ParseBool(args[5])
			if err != nil {
				return err
			}

			ticketData := types.TicketData{} // TODO: change me to correct data
			eventData := types.EventDetails{}

			msg := em.NewMsgCreateEvent(args[0], num, args[3], cliCtx.GetFromAddress(), boo, ticketData, eventData)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

//  NewMsgNewOwner(eventName string, previousOwnerAddress sdk.AccAddress, newOwnerAddress sdk.AccAddress, newOwner string)
func GetCmdNewOwner(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "setNewOwner [eventName] [newOwnerAddress] [newOwnerName]",
		Short: "set a new owner",
		Long:  "Change the owner of the event",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := em.NewMsgNewOwner(args[0], cliCtx.GetFromAddress(), addr, args[2])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdCloseEvent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "closeEvent [eventID] [eventOwnerAddress]",
		Short: "Close an event",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			eventID := args[0]

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}
			msg := em.NewMsgCloseEvent(eventID, sdk.AccAddress(args[1]))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			cliCtx.PrintResponse = true
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
