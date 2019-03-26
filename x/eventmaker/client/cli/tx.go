package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	eventTypes "github.com/marbar3778/tic_mark/types"
	"github.com/marbar3778/tic_mark/x/eventmaker"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

// InitialPrice  sdk.Coin `json:"ticket_price"`
// MarkUpAllowed int      `json:"mark_up_allowed"`
// TotalTickets  int      `json:"total_tickets"`
// TicketsSold   int      `json:"tickets_sold"`
// Resale        bool     `json:"resale"`
// }
// type EventDetails struct {
// LocationName string `json:"location_name"`
// Address      string `json:"address"` // Address
// City         string `json:"city"`    // City in
// Country      string `json:"country"` // Country
// Date         string `json:"date"`    // date of
// }

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

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

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

			ticketData := eventTypes.TicketData{}
			eventData := eventTypes.EventDetails{}

			msg := eventmaker.NewMsgCreateEvent(args[0], num, args[3], cliCtx.GetFromAddress(), boo, ticketData, eventData)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
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

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := eventmaker.NewMsgNewOwner(args[0], cliCtx.GetFromAddress(), addr, args[2])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}
