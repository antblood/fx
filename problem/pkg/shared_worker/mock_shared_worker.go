package shared_worker

import (
	"fmt"

	"go.uber.org/fx"
)

var MockSharedWorkerModule = fx.Options(
	fx.Provide(NewMockSharedWorker),
	fx.Provide(RegisterActivities),
)

func NewMockSharedWorker() SharedWorkerOuts {
	return SharedWorkerOuts{
		Worker: &MockSharedWorker{},
	}
}

// Internal struct, can't change this
type MockSharedWorker struct {
	Activities []string
}

func (w *MockSharedWorker) RegisterActivity(activityName string) {
	fmt.Println("I'm inside mock shared worker:", activityName)
	w.Activities = append(w.Activities, activityName)
}
