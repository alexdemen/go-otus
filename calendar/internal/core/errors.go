package core

import "errors"

var ErrEventNotExist = errors.New("event not exist")

var ErrNoEvents = errors.New("no events")

var ErrDateBusy = errors.New("this time is busy")
