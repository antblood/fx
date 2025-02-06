package activities_consumer

import (
	"fmt"

	"go.uber.org/fx"
)

type ActivitiesConsumerParams struct {
	fx.In

	ActivitiesList []string `optional:"true"`
}

var Module = fx.Module("activities_consumer",
	fx.Invoke(NewActivitiesConsumer),
)

func NewActivitiesConsumer(params ActivitiesConsumerParams) bool {
	if params.ActivitiesList == nil {
		fmt.Println("No activities found :(")
		return false
	}

	for _, activity := range params.ActivitiesList {
		fmt.Println("Consumed activity:", activity)
	}

	return true
}
