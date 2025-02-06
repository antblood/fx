package main

import (
	"go.uber.org/fx"
	"neeraj.com/fx/problem/pkg/shared_worker"
)

func opts() fx.Option {
	return fx.Options(
		shared_worker.SharedWorkerModule,
	)
}

func main() {
	fx.New(opts()).Run()
}
