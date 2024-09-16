package valueobject

import (
	"chater/internal/domain/validation"
)

type ChatName struct {
	value string
}

func NewChatName(chatName string) (*ChatName, error) {
	// validate chat name
	err := validation.ValidateChatName(chatName)
	if err != nil {
		return nil, err
	}
	return &ChatName{value: chatName}, nil
}

func (cn *ChatName) Value() string {
	return cn.value
}
