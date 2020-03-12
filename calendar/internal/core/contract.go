package core

import (
	"context"
	"time"
)

type Store interface {
	Add(cxt context.Context, event Event) (Event, error)
	Edit(cxt context.Context, event Event) error
	Remove(cxt context.Context, event Event) error
	List(cxt context.Context) ([]Event, error)
}

type Event struct {
	Id          int64
	Name        string
	Description *string
	StartDate   time.Time
	Duration    time.Duration
}
