package market

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	mmtypes "github.com/marbar3778/tic_mark/x/market/types"
)

type GenesisState struct {
	EventTickets []mmtypes.Ticket `json:"event_tickets"`
	MarketPlace  []mmtypes.Ticket `json:"market_place"`
}

func NewGenesisState() GenesisState {
	return GenesisState{
		EventTickets: nil,
		ClosedEvents: nil,
	}
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	for _, tickets := range data.EventTickets {
		// k.SetEvent(ctx, record.EventID, record, k.eKey)
		// when setting the tickets, check if key is already there, if not create new, else append and set
	}
	for _, tickets := range data.MarketPlace {
		// k.SetEvent(ctx, record.EventID, record, k.ceKey)
	}
}

// func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
// 	var openRecords []mmtypes.Event
// 	openIterator := k.GetAllEvents(ctx, k.eKey)
// 	for ; openIterator.Valid(); openIterator.Next() {
// 		key := string(openIterator.Key())
// 		var oRecord mmtypes.Event
// 		oRecord, ok := k.GetOpenEvent(ctx, key)
// 		if !ok {
// 			fmt.Println("Error, key: %s does not have a value", key)
// 		}
// 		openRecords = append(openRecords, oRecord)
// 	}
// 	var closedRecords []mmtypes.Event
// 	closedIterator := k.GetAllEvents(ctx, k.ceKey)
// 	for ; closedIterator.Valid(); closedIterator.Next() {
// 		key := string(closedIterator.Key())
// 		var cRecord mmtypes.Event
// 		cRecord, ok := k.GetOpenEvent(ctx, key)
// 		if !ok {
// 			fmt.Println("Error, key: %s does not have a value", key)
// 		}
// 		closedRecords = append(closedRecords, cRecord)
// 	}

// 	return GenesisState{
// 		OpenEvents:   openRecords,
// 		ClosedEvents: closedRecords,
// 	}
// }

// func ValidateGenesis(data GenesisState) error {
// 	for _, data := range data.EventRecords {
// 		if data.EventOwner == nil {
// 			return fmt.Errorf("Event needs an owner, current owner: %s", data.EventOwner)
// 		}
// 		if data.EventName == nil {
// 			return fmt.Errorf("Event must have a name, current name: %s", data.EventName)
// 		}
// 		if data.TicketData == nil {
// 			return fmt.Errorf("Invalid ticketData")
// 		}
// 	}
// }
