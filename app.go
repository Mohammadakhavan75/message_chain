package app

import (
	"io"

	// Tendermint and Cosmos SDK packages
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	// Import your custom module packages
	"github.com/Mohammadakhavan75/message_chain/x/message"
	messagekeeper "github.com/Mohammadakhavan75/message_chain/x/message/keeper"
	messagetypes "github.com/Mohammadakhavan75/message_chain/x/message/types"
)

// DummyApp defines your custom application.
type DummyApp struct {
	*baseapp.BaseApp
	cdc  *codec.LegacyAmino
	keys map[string]*sdk.KVStoreKey

	// Custom module keeper: handles storing and retrieving messages.
	MessageKeeper messagekeeper.Keeper
}

// NewDummyApp creates a new instance of DummyApp.
func NewDummyApp(logger log.Logger, db dbm.DB, traceStore io.Writer) *DummyApp {
	// Create a codec for encoding/decoding.
	cdc := codec.NewLegacyAmino()

	// Initialize the BaseApp with a name, logger, database, and codec.
	bApp := baseapp.NewBaseApp("dummychain", logger, db, cdc)

	// Define and create store keys. Here, "message" will be used for our custom module.
	keys := sdk.NewKVStoreKeys("message")

	// Instantiate the app structure.
	app := &DummyApp{
		BaseApp: bApp,
		cdc:     cdc,
		keys:    keys,
	}

	// Initialize the custom module keeper.
	// It requires the codec and the KVStoreKey for "message".
	app.MessageKeeper = messagekeeper.NewKeeper(app.cdc, keys["message"])

	// ----------------------------------------------------
	// **Wire in the Message Module**
	// ----------------------------------------------------

	// 1. Register the module's message route:
	// This tells BaseApp how to handle messages of type MsgStoreMessage.
	app.Router().AddRoute(messagetypes.RouterKey, message.NewHandler(app.MessageKeeper))

	// 2. Register the module's querier route:
	// This makes it possible to query your module’s state (e.g. retrieve all messages).
	app.QueryRouter().AddRoute(messagetypes.RouterKey, message.NewQuerier(app.MessageKeeper))

	// ----------------------------------------------------
	// Mount your KVStores to the BaseApp:
	// This will allow the BaseApp to persist state for your module.
	for _, key := range keys {
		app.MountStore(key, sdk.StoreTypeIAVL)
	}

	// Load the latest state from the store.
	if err := app.LoadLatestVersion(keys["message"]); err != nil {
		panic(err)
	}

	return app
}
