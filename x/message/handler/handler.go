package message

import (
	"github.com/Mohammadakhavan75/message_chain/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k messagekeeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case *types.MsgStoreMessage:
			return handleMsgStoreMessage(ctx, k, *msg)
		default:
			errMsg := "unrecognized message type"
			return nil, sdk.ErrUnknownRequest(errMsg)
		}
	}
}

func handleMsgStoreMessage(ctx sdk.Context, k messagekeeper.Keeper, msg types.MsgStoreMessage) (*sdk.Result, error) {
	k.StoreMessage(ctx, msg)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.RouterKey,
			sdk.NewAttribute("sender", msg.Sender.String()),
			sdk.NewAttribute("content", msg.Content),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
