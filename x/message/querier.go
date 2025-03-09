package keeper

import (
	"encoding/json"

	// "cosmossdk.io/api/tendermint/abci"
	"github.com/Mohammadakhavan75/message_chain/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/abci"
)

func NewQuerier(k types.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case "all":
			messages := k.GetAllMessages(ctx)
			res, err := json.Marshal(messages)
			if err != nil {
				return nil, err
			}
			return res, nil
		default:
			return nil, sdk.ErrUnknownRequest("unknown query endpoint")
		}
	}
}
