package scheduler

import (
	"context"
	"github.com/alexdemen/go-otus/calendar/internal/core"
	"testing"
	"time"
)

type TestExplorer struct {
}

func (t TestExplorer) Get(ctx context.Context) ([]core.Event, error) {
	res := make([]core.Event, 0)
	res = append(res, core.Event{
		Id:          1,
		Name:        "1212",
		Description: nil,
		StartDate:   time.Now(),
		Duration:    666,
	})

	return res, nil
}

func TestNewScheduler(t *testing.T) {

	sch := NewScheduler(TestExplorer{}, "amqp://guest:guest@localhost:5672/")
	sch.run(context.Background())

	time.Sleep(5 * time.Minute)
}
