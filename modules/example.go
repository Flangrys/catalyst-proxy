package modules

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/flangrys/catalyst-proxy/services"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Define your field for your workload.
type ExampleWorker struct {
	Id string
}

// Define your logic for the routine invocation.
func (ptr *ExampleWorker) Invoke(ctx context.Context, wg *sync.WaitGroup, ch chan services.ServiceChannel) {
	var (
		id  uuid.UUID
		err error
	)

	// Try gen a random uuid for this worker.
	if id, err = uuid.NewRandom(); err == nil {

		// Send a message with the reason of the invocation failure.
		ch <- services.ServiceChannel{
			Reason: errors.New("cannot generate a random id"),
			Failed: true,
		}

		return
	}

	ptr.Id = id.String()

	defer wg.Done()

	select {
	case <-ctx.Done():
		logrus.Info("Reciving SIGTERM signal.")

		ch <- services.ServiceChannel{
			Failed: false,
			Reason: nil,
		}

	default:
		logrus.Info("Normal execution stop")
		time.Sleep(time.Duration(20) * time.Second)
	}
}
