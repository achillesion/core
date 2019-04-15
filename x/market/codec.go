package market

import "github.com/cosmos/cosmos-sdk/codec"

// RegisterCodec : Register codec msgs
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateTicket{}, "market/CreateTicket", nil)
	cdc.RegisterConcrete(MsgResaleTicket{}, "market/ResaleTicket", nil)
}
