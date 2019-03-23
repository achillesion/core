package app

import (
	"github.com/tendermint/tendermint/libs/log"

	"github.com/abcInfinity/tic_mark/x/eventmaker"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
)

const (
	appName = "tic_mark"
)

type EventMarketApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	keyMain          *sdk.KVStoreKey
	keyAccount       *sdk.KVStoreKey
	keyEM            *sdk.KVStoreKey // key for upcoming event store
	keyECM           *sdk.KVStoreKey // key for closed event store
	keyFeeCollection *sdk.KVStoreKey
	keyParams        *sdk.KVStoreKey
	tkeyParams       *sdk.TransientStoreKey

	accountKeeper       auth.AccountKeeper
	bankKeeper          bank.Keeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	paramsKeeper        params.Keeper
	tmKeeper            eventmaker.Keeper
}

func NewEventMarketApp(logger log.Logger, db dbm.DB) *EventMarketApp {
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &EventMarketApp{
		BaseApp:          bApp,
		cdc:              cdc,
		keyMain:          sdk.NewKVStoreKey("main"),
		keyAccount:       sdk.NewKVStoreKey("acc"),
		keyEM:            sdk.NewKVStoreKey("em"),
		keyECM:           sdk.NewKVStoreKey("closed_events"),
		keyFeeCollection: sdk.NewKVStoreKey("fee_collection"),
		keyParams:        sdk.NewKVStoreKey("params"),
		tkeyParams:       sdk.NewTransientStoreKey("transient_params"),
	}

	app.paramsKeeper = params.NewKeeper(app.cdc, app.keyParams, app.tkeyParams)

	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount,
		app.paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)

	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		app.paramsKeeper.Subspace(bank.DefaultParamspace),
		bank.DefaultCodespace,
	)

	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(cdc, app.keyFeeCollection)

	app.tmKeeper = eventmaker.NewKeeper(
		app.bankKeeper,
		app.keyEM,
		app.keyECM,
		app.cdc,
	)

	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))

	app.Router().
		AddRoute("bank", bank.NewHandler(app.bankKeeper)).
		AddRoute("eventmaker", eventmaker.NewHandler(app.tmKeeper))

	app.QueryRouter().
		AddRoute("eventmaker", eventmaker.NewQuerier(app.tmKeeper))

	app.SetInitChainer(app.initChain)

	app.MountStores(
		app.keyMain,
		app.keyAccount,
		app.keyEM,
		app.keyECM,
		app.keyFeeCollection,
		app.keyParams,
		app.tkeyParams,
	)

	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}

	return app
}

func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	bank.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	eventmaker.RegisterCodec(cdc)
	return cdc
}

type GenesisState struct {
	AuthData auth.GenesisState   `json:"auth"`
	BankData bank.GenesisState   `json:"bank"`
	Accounts []*auth.BaseAccount `json:"accounts"`
}

func (app *EventMarketApp) initChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	stateJSON := req.AppStateBytes

	genesisState := new(GenesisState)
	err := app.cdc.UnmarshalJSON(stateJSON, genesisState)
	if err != nil {
		panic(err)
	}

	for _, acc := range genesisState.Accounts {
		acc.AccountNumber = app.accountKeeper.GetNextAccountNumber(ctx)
		app.accountKeeper.SetAccount(ctx, acc)
	}

	auth.InitGenesis(ctx, app.accountKeeper, app.feeCollectionKeeper, genesisState.AuthData)
	bank.InitGenesis(ctx, app.bankKeeper, genesisState.BankData)

	return abci.ResponseInitChain{}
}
