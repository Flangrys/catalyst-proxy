package services

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"slices"
	"sync"
	"syscall"

	"github.com/flangrys/catalyst-proxy/config"
	"github.com/sirupsen/logrus"
)

type Scheduler struct {
	Config    *config.Configuration
	Workers   []Worker
	IsRunning bool

	ctx        context.Context
	cancelCtx  context.CancelFunc
	waitGroup  sync.WaitGroup
	osSignalCh chan os.Signal
}

var (
	scheduler Scheduler
)

func NewWithConfig(conf *config.Configuration) (*Scheduler, error) {
	logrus.Info("Initializing Scheduler with configuration.")

	// Avoid multiple invocation at the same program.
	if scheduler.IsRunning {
		return nil, ErrIlegalServiceInit
	}

	// Setup scheduler fields.
	scheduler.ctx, scheduler.cancelCtx = context.WithCancel(context.Background())
	scheduler.osSignalCh = make(chan os.Signal, 1)
	scheduler.IsRunning = true
	scheduler.Config = conf

	return &scheduler, nil
}

// Init the lifecycle manager.
func (s *Scheduler) InitLifecycleManager() {
	if !s.IsRunning {
		logrus.Fatal("Cannot setup the lifecycle manager when the scheduler is not running.")
	}

	// Capture the signal sent through the channel 's.osSignalCh' when the
	// program must to finish.
	signal.Notify(s.osSignalCh, syscall.SIGINT, syscall.SIGTERM)

	logrus.Info("Waiting for closing connections...")
	s.IsRunning = false
	s.cancelCtx()

}

// Waits for all scheduled task to be completed.
func (s *Scheduler) WaitAll() {
	s.waitGroup.Wait()
	// Close the service channel.
}

// Append a new task to the task queue.
func (s *Scheduler) AddTask(task *Task) error {
	if task == nil {
		return ErrNullPointer
	}

	// Add the task to the scheduler tasks queue.

	return nil
}

// Invoke an existent task with their id.
func (s *Scheduler) RunTask(id string) error {

	// Find the task by their id.
	found := slices.IndexFunc(s.Workers, func(worker Worker) bool {

		// Asserts when `Task` implements `Worker`
		if wk, ok := worker.(Task); ok {
			return wk.Id == id
		}

		return false
	})

	if found == -1 {
		return fmt.Errorf("task with the id: %s were not found", id)
	}

	routine := s.Workers[found]

	// Setup channel.
	serviceChannel := make(chan ServiceChannel)

	// Magic happends here.
	s.waitGroup.Add(1)

	go routine.Invoke(s.ctx, &s.waitGroup, serviceChannel)

	// When the task ends, remove from the task queue and clean the channel.

	return nil
}
