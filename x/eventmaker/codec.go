package eventmaker

import "github.com/cosmos/cosmos-sdk/codec"

// RegisterCodec : Register codec msgs
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateEvent{}, "eventmaker/CreateEvent", nil)
	cdc.RegisterConcrete(MsgNewOwner{}, "eventmaker/NewOwner", nil)
	cdc.RegisterConcrete(MsgCloseEvent{}, "eventmaker/CloseEvent", nil)
}

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
