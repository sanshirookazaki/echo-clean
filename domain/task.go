package domain

// Task .
type Task struct {
	ID     int
	Userid int
	Task   string
	Status int
}

// Tasks .
type Tasks []Task
