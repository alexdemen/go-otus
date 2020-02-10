package core

import "errors"

var (
	ErrEventNotExist = errors.New("event not exist")
	ErrNoEvents      = errors.New("no events")
	ErrDateBusy      = errors.New("this time is busy")
)
