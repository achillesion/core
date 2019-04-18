package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	emd "github.com/marbar3778/tic_mark/x/eventmaker/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

type ModuleClient struct {
	storekey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	ticketQueryCmd := &cobra.Command{
		Use:   "eventmaker",
		Short: "Querying commands for the eventmaker module",
	}
	ticketQueryCmd.AddCommand(client.GetCommands(
		emd.GetCmdGetOpenEvent(mc.storekey, mc.cdc),
		emd.GetCmdGetClosedEvent(mc.storekey, mc.cdc),
		emd.GetCmdGetOwner(mc.storekey, mc.cdc)
	)...)

	return ticketQueryCmd

}

func (mc ModuleClient) GetTxCmd() *cobra.Command {
	ticketTxCmd := &cobra.Command{
		Use:   "eventmaker",
		Short: "eventmaker tx subcommands",
	}

	ticketTxCmd.AddCommand(client.PostCommands(
		emd.GetCmdCreateEvent(mc.cdc),
		emd.GetCmdNewOwner(mc.cdc),
		emd.GetCmdCloseEvent(mc.cdc)
	)...)

	return ticketTxCmd
}
