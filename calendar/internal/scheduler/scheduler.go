package scheduler

import (
	"context"
	"github.com/alexdemen/go-otus/calendar/internal/core"
	"time"
)

func ScheduleEvent(cxt context.Context, store core.Explorer) {
	tikChan := time.Tick(30 * time.Minute)

	select {
	case <-tikChan:
		discover(cxt, store)
	case <-cxt.Done():
		return

	}
}

func discover(cxt context.Context, exp core.Explorer) error {
	events, err := exp.Get(cxt)
	if err != nil {
		return err
	}

	if len(events) == 0 {
		return nil
	}

	return nil
}
