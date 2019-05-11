package main

import (
	"sync"

	"github.com/arifsetiawan/camunda-external-task-client/camunda"
	"github.com/arifsetiawan/camunda-external-task-client/handlers"

	"github.com/robfig/cron"
	"github.com/rs/zerolog/log"
)

func main() {
	var wg sync.WaitGroup

	c := cron.New()

	client := &camunda.Client{
		APIURL: "http://localhost:8080",
	}

	availableLeaveDays := &camunda.FetchAndLock{
		WorkerID: "worker1",
		MaxTasks: 1,
		Topics: []camunda.Topic{
			camunda.Topic{
				Name:         "available-leave-days",
				LockDuration: 10000,
			},
		},
	}

	email := &camunda.FetchAndLock{
		WorkerID: "worker2",
		MaxTasks: 1,
		Topics: []camunda.Topic{
			camunda.Topic{
				Name:         "send-rejection-email",
				LockDuration: 10000,
			},
			camunda.Topic{
				Name:         "send-approval-email",
				LockDuration: 10000,
			},
		},
	}

	c.AddFunc("@every 5s", call(client, availableLeaveDays, handlers.AvailableLeaveDaysHandler))

	c.AddFunc("@every 5s", call(client, email, handlers.EmailHandler))

	c.Start()

	wg.Add(1)
	wg.Wait()
}

func call(client *camunda.Client, param *camunda.FetchAndLock, handler func(client *camunda.Client, tasks []camunda.Task) error) func() {
	return func() {
		tasks, err := client.FetchAndLock(param)
		if err != nil {
			log.Error().Err(err).Msg("fetch and lock task")
			return
		}

		log.Info().Interface("topics", param.Topics).Interface("tasks", tasks).Msg("Fetch and lock")

		if len(tasks) > 0 {
			err = handler(client, tasks)
			if err != nil {
				log.Error().Err(err).Msg("task handler")
				return
			}
		}
	}
}
