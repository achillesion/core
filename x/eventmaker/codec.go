package eventmaker

import "github.com/cosmos/cosmos-sdk/codec"

// RegisterCodec : Register codec msgs
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateEvent{}, "eventmaker/CreateEvent", nil)
	cdc.RegisterConcrete(MsgNewOwner{}, "eventmaker/NewOwner", nil)
	cdc.RegisterConcrete(MsgCloseEvent{}, "eventmaker/CloseEvent", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
