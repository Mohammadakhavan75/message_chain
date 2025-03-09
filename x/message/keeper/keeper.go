package messagekeeper

import (
	"github.com/Mohammadakhavan75/message_chain/x/message/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	cdc *codec.LegacyAmino
	key *sdk.KVStoreKey
}

func NewKeeper(cdc *codec.LegacyAmino, key *sdk.KVStoreKey) types.Keeper {
	return Keeper{cdc: cdc, key: key}
}

func (k Keeper) StoreMessage(ctx sdk.Context, msg types.MsgStoreMessage) {
	store := ctx.KVStore(k.key)
	// Create a unique key, e.g., using current block height and sender address
	key := append([]byte("message-"), []byte(ctx.BlockHeader().Height.String()+"-"+msg.Sender.String())...)
	store.Set(key, []byte(msg.Content))
}

func (k Keeper) GetAllMessages(ctx sdk.Context) []string {
	store := ctx.KVStore(k.key)
	iterator := sdk.KVStorePrefixIterator(store, []byte("message-"))
	defer iterator.Close()

	var messages []string
	for ; iterator.Valid(); iterator.Next() {
		messages = append(messages, string(iterator.Value()))
	}
	return messages
}
