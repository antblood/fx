package shared_worker

import "go.uber.org/fx"

type IWorker interface {
	RegisterActivity(activityName string)
}

type WorkerParams struct {
	fx.In

	Worker IWorker
}

type RegisterActivitiesOut struct {
	fx.Out

	Activities []string
}

func RegisterActivities(params WorkerParams) RegisterActivitiesOut {
	params.Worker.RegisterActivity("activity1")
	params.Worker.RegisterActivity("activity2")

	return RegisterActivitiesOut{
		Activities: []string{},
	}
}
