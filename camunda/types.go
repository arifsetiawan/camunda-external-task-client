package camunda

// Topic ...
type Topic struct {
	Name         string `json:"topicName"`
	LockDuration int    `json:"lockDuration"`
}

// FetchAndLock ...
type FetchAndLock struct {
	WorkerID string  `json:"workerId"`
	MaxTasks int     `json:"maxTasks"`
	Topics   []Topic `json:"topics"`
}

// Task ...
type Task struct {
	ID        string              `json:"id"`
	TopicName string              `json:"topicName"`
	Variables map[string]Variable `json:"variables"`
}

// Variable ...
type Variable struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// Complete ...
type Complete struct {
	WorkerID  string              `json:"workerId"`
	Variables map[string]Variable `json:"variables"`
}
