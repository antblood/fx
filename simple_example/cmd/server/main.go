package main

import (
	"go.uber.org/fx"
	"neeraj.com/fx/simple_example/pkg/message_constructor"
	"neeraj.com/fx/simple_example/pkg/message_consumer"
)

func opts() fx.Option {
	return fx.Options(
		message_constructor.Module,
		message_consumer.Module,
	)
}

func main() {
	fx.New(opts()).Run()
}
