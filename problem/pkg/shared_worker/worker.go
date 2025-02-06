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

	// ===================
	// I want you to keep this part of the code as is
	params.Worker.RegisterActivity("activity1")
	params.Worker.RegisterActivity("activity2")
	// till here
	// ===================

	// if you just add "activity1" and "activity2" here, it will work
	// but if I add more activities, it will break
	// because new activities will need to be added here manually
	return RegisterActivitiesOut{
		Activities: []string{},
	}
}
