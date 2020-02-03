package store

import "time"

type Store interface {
	Add(event Event) error
	Edit(event Event) error
	Remove(event Event) error
	List() ([]Event, error)
}

type Event struct {
	Id          int64
	Name        string
	Description string
	StartDate   time.Time
	FinishDate  time.Time
}
