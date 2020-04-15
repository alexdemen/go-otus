package logger

import (
	"context"
	"github.com/alexdemen/go-otus/calendar/pkg/calendarpb"
)

type GRPCLogger struct {
	next calendarpb.EventServiceServer
}

func (G GRPCLogger) Save(cxt context.Context, event *calendarpb.Event) (*calendarpb.Event, error) {
	res, err := G.next.Save(cxt, event)

	return res, err
}

func (G GRPCLogger) Remove(cxt context.Context, event *calendarpb.Event) (*calendarpb.Result, error) {
	res, err := G.next.Remove(cxt, event)

	return res, err
}

func (G GRPCLogger) List(cxt context.Context, user *calendarpb.User) (*calendarpb.EventList, error) {
	res, err := G.next.List(cxt, user)

	return res, err
}

func NewMiddlewareLogger(next calendarpb.EventServiceServer) calendarpb.EventServiceServer {
	var server calendarpb.EventServiceServer = GRPCLogger{
		next: next,
	}

	return server
}
