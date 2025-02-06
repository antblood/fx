package message_constructor

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewMessage),
)

type Message string

type Constructor_type struct {
	fx.Out

	Message Message
}

func NewMessage() Constructor_type {
	return Constructor_type{
		Message: Message("0123456789"),
	}
}
