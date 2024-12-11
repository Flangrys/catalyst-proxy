package services

import (
	"context"
	"sync"
)

// Represents a task interface.
type Worker interface {
	Invoke(context.Context, *sync.WaitGroup, chan ServiceChannel)
}

// Represents a task.
type Task struct {
	Id string

	Worker
}
