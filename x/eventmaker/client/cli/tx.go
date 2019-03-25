package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/marbar3778/tic_mark/x/eventmaker"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

//  NewMsgCreateEvent(eventName string, totalTickets int, ticketsSold int, eventOwner string, eventOwnerAddress sdk.AccAddress, resale bool)
func GetCmdCreateEvent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createEvent [eventName] [totalTickets] [ticketsSold] [eventOwner] [resale]",
		Short: "Create Event",
		Long:  "Create a event",
		Args:  cobra.ExactArgs(5),
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

			num2, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}

			boo, err := strconv.ParseBool(args[5])
			if err != nil {
				return err
			}

			msg := eventmaker.NewMsgCreateEvent(args[0], num, num2, args[3], cliCtx.GetFromAddress(), boo)
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
		RunE: func(cmd *codec.Codec, args []string) error {
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
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}
