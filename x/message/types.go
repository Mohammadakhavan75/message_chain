package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = "message"

// MsgStoreMessage defines a message for storing a text message.
type MsgStoreMessage struct {
	Sender  sdk.AccAddress `json:"sender" yaml:"sender"`
	Content string         `json:"content" yaml:"content"`
}

func NewMsgStoreMessage(sender sdk.AccAddress, content string) *MsgStoreMessage {
	return &MsgStoreMessage{Sender: sender, Content: content}
}

func (msg MsgStoreMessage) Route() string { return RouterKey }
func (msg MsgStoreMessage) Type() string  { return "store_message" }
func (msg MsgStoreMessage) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress("missing sender address")
	}
	if len(msg.Content) == 0 {
		return sdk.ErrUnknownRequest("content cannot be empty")
	}
	return nil
}
func (msg MsgStoreMessage) GetSignBytes() []byte {
	// Use amino JSON encoding for signing
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}
func (msg MsgStoreMessage) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
