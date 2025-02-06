package shared_worker

import (
	"fmt"

	"go.uber.org/fx"
)

var SharedWorkerModule = fx.Options(
	fx.Provide(NewSharedWorker),
	fx.Invoke(RegisterActivities),
)

type SharedWorkerOuts struct {
	fx.Out

	Worker IWorker
}

func NewSharedWorker() SharedWorkerOuts {
	return SharedWorkerOuts{
		Worker: &SharedWorker{},
	}
}

type SharedWorker struct{}

func (w *SharedWorker) RegisterActivity(activityName string) {
	fmt.Println("Registering activity:", activityName)
}
