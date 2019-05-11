package handlers

import (
	"time"

	"github.com/arifsetiawan/camunda-external-task-client/camunda"
)

// AvailableLeaveDaysHandler ...
func AvailableLeaveDaysHandler(client *camunda.Client, tasks []camunda.Task) error {

	for _, task := range tasks {

		// pretend to do something
		time.Sleep(2 * time.Second)

		// complete
		err := client.Complete(task.ID, &camunda.Complete{
			WorkerID: "worker1",
			Variables: map[string]camunda.Variable{
				"daysAvailable": camunda.Variable{
					Value: true,
				},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
