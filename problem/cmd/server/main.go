package main

import (
	"go.uber.org/fx"
)

func opts() fx.Option {
	return fx.Options()
}

func main() {
	fx.New(opts()).Run()
}
