package message_consumer

import (
	"fmt"

	"go.uber.org/fx"
	"neeraj.com/fx/simple_example/pkg/message_constructor"
)

var Module = fx.Options(
	fx.Invoke(PrintMessageLength),
)

type ins struct {
	fx.In

	Message message_constructor.Message
}

// Another function that also depends on the Message
func PrintMessageLength(ins ins) {
	fmt.Printf("Message length from inputs: %d\n", len(ins.Message))
}
