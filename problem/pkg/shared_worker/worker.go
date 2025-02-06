package shared_worker

import "go.uber.org/fx"

type IWorker interface {
	RegisterActivity(activityName string)
}

type WorkerParams struct {
	fx.In

	Worker IWorker
}

func RegisterActivities(params WorkerParams) {
	params.Worker.RegisterActivity("activity1")
	params.Worker.RegisterActivity("activity2")
}
