package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
	market "github.com/marbar3778/tic_mark/x/market/client/cli"


)
type ModuleClient struct {
	storeKey string
	cdc *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	ticketQueryCmd := &cobra.Command{
		Use: "ticketmarket",
		Short: "Query for the ticket market"
	}

	ticketQueryCmd.AddCommand(client.GetCommand(
		market.GetCmdGetTickets(mc.storeKey, mc.cdc),
		market.GetCmdGetTicket(mc.storeKey, mc.cdc)
	)...)

	ticketQueryCmd.AddCommand(client.PostCommands(
		market.GetCmdCreateTicket(mc.cdc),
		market.GetCmdResellTicket(mc.cdc),
	)...)
	return ticketQueryCmd
}