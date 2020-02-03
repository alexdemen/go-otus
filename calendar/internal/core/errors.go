package core

import "errors"

type ErrEventNotExist struct{}

type ErrNoEvents struct{}

func (n ErrNoEvents) Error() string {
	return "No events."
}

func (e ErrEventNotExist) Error() string {
	return "Event not exist."
}

var ErrDateBusy = errors.New("this time is busy")
