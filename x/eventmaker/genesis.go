package eventmaker

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	emtypes "github.com/marbar3778/tic_mark/x/eventmaker/types"
)

type GenesisState struct {
	OpenEvents   []emtypes.Event `json:"open_event"`
	ClosedEvents []emtypes.Event `json:"closed_event"`
}

func NewGenesisState() GenesisState {
	return GenesisState{
		OpenEvents:   nil,
		ClosedEvents: nil,
	}
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, k BaseKeeper, data GenesisState) {
	for _, record := range data.OpenEvents {
		k.SetEvent(ctx, record.EventID, record, k.eKey)
	}
	for _, record := range data.ClosedEvents {
		k.SetEvent(ctx, record.EventID, record, k.ceKey)
	}
}

func ExportGenesis(ctx sdk.Context, k BaseKeeper) GenesisState {
	var openRecords []emtypes.Event
	openIterator := k.GetAllEvents(ctx, k.eKey)
	for ; openIterator.Valid(); openIterator.Next() {
		key := string(openIterator.Key())
		var oRecord emtypes.Event
		oRecord, ok := k.GetOpenEvent(ctx, key)
		if !ok {
			fmt.Println("Error, key: %s does not have a value", key)
		}
		openRecords = append(openRecords, oRecord)
	}
	var closedRecords []emtypes.Event
	closedIterator := k.GetAllEvents(ctx, k.ceKey)
	for ; closedIterator.Valid(); closedIterator.Next() {
		key := string(closedIterator.Key())
		var cRecord emtypes.Event
		cRecord, ok := k.GetOpenEvent(ctx, key)
		if !ok {
			fmt.Println("Error, key: %s does not have a value", key)
		}
		closedRecords = append(closedRecords, cRecord)
	}

	return GenesisState{
		OpenEvents:   openRecords,
		ClosedEvents: closedRecords,
	}
}

func ValidateGenesis(data GenesisState) error {
	for _, data := range data.EventRecords {
		if data.EventOwner == nil {
			return fmt.Errorf("Event needs an owner, current owner: %s", data.EventOwner)
		}
		if data.EventName == nil {
			return fmt.Errorf("Event must have a name, current name: %s", data.EventName)
		}
		if data.TicketData == nil {
			return fmt.Errorf("Invalid ticketData")
		}
	}
}
