package eventmaker

import (
	"encoding/json"

 	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

 	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ sdk.AppModule      = AppModule{}
	_ sdk.AppModuleBasic = AppModuleBasic{}
)

const ModuleName = "eventmaker"

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return ModuleName
}

 func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

 // Validation check of the Genesis
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	// once json successfully marshalled, passes along to genesis.go
	return ValidateGenesis(data)
}

type AppModule struct {
	AppModuleBasic
	keeper     BaseKeeper
}

func NewAppModule(k BaseKeeper, bankKeeper bank.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

func (AppModule) Name() string {
	return ModuleName
}

func (em AppModule) RegisterInvariants(ir sdk.InvariantRouter) {}

func (em AppModule) Route() string {
	return RouterKey
}

func (em AppModule) NewHandler() types.Handler {
	return NewHandler(em.keeper)
}
func (em AppModule) QuerierRoute() string {
	return ModuleName
}

func (em AppModule) NewQuerierHandler() types.Querier {
	return NewQuerier(em.keeper)
}

func (em AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) types.Tags {
	return sdk.EmptyTags()
}

func (em AppModule) EndBlock(types.Context, abci.RequestEndBlock) ([]abci.ValidatorUpdate, types.Tags) {
	return []abci.ValidatorUpdate{}, sdk.EmptyTags()
}

func (em AppModule) InitGenesis(ctx types.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	return InitGenesis(ctx, em.keeper, genesisState)
}

func (em AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, em.keeper)
	return ModuleCdc.MustMarshalJSON(gs)
}