
# Camunda External Task Client

Example how to handle Camunda external task with Golang

See https://medium.com/@arifsetiawan/my-journey-with-camunda-toc-3030da004511

## How it works

1. Run cron to fetch tasks

```
c := cron.New()

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

c.AddFunc("@every 5s", call(client, availableLeaveDays, handlers.AvailableLeaveDaysHandler))

c.Start()
```

2. Task handler

```
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
```

## Run

Make sure Camunda is running

```
go run main.go
```