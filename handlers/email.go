package handlers

import (
	"time"

	"github.com/arifsetiawan/camunda-external-task-client/camunda"
)

// EmailHandler ...
func EmailHandler(client *camunda.Client, tasks []camunda.Task) error {

	for _, task := range tasks {

		// pretend to do something
		time.Sleep(2 * time.Second)

		// complete
		err := client.Complete(task.ID, &camunda.Complete{
			WorkerID: "worker2",
		})
		if err != nil {
			return err
		}
	}

	return nil
}
