package service

import (
	"context"
	"github.com/alexdemen/go-otus/calendar/internal/calendarpb"
	"github.com/alexdemen/go-otus/calendar/internal/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EventServer struct {
	store core.Store
}

func NewEventServer(store *core.Store) *EventServer {
	return &EventServer{
		store: *store,
	}
}

func (e EventServer) Save(context.Context, *calendarpb.Event) (*calendarpb.Event, error) {
	panic("implement me")
}

func (e EventServer) Remove(context.Context, *calendarpb.Event) (*calendarpb.Result, error) {
	panic("implement me")
}

func (e EventServer) List(context.Context, *calendarpb.User) (*calendarpb.EventList, error) {
	list, err := e.store.List()
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	panic("implement me")
}
