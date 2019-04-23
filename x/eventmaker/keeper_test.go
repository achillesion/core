package eventmaker

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/params"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

type testInput struct {
	cdc *codec.Codec
	ctx sdk.Context
	ak  auth.AccountKeeper
	pk  params.Keeper
}

func setupTestInput() testInput {
	db := dbm.NewMemDB()

	cdc := codec.New()
	auth.RegisterBaseAccount(cdc)

	keyAccount := sdk.NewKVStoreKey("acc")
	// keyEM := sdk.NewKVStoreKey("em")
	// keyECM := sdk.NewKVStoreKey("closed_events")
	keyParams := sdk.NewKVStoreKey("params")
	tkeyParams := sdk.NewTransientStoreKey("transient_params")

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyAccount, sdk.StoreTypeIAVL, db)
	// ms.MountStoreWithDB(keyEM, sdk.StoreTypeIAVL, db)
	// ms.MountStoreWithDB(keyECM, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)

	pk := params.NewKeeper(cdc, keyParams, tkeyParams)
	ak := auth.NewAccountKeeper(cdc, keyAccount, pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)

	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	ak.SetParams(ctx, auth.DefaultParams())

	return testInput{cdc: cdc, ctx: ctx, ak: ak, pk: pk}
}

func TestKeeper(t *testing.T) {
	input := setupTestInput()
	ctx := input.ctx
	eventMaker := NewKeeper()
}
