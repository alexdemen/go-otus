package service

import (
	"context"
	"github.com/alexdemen/go-otus/calendar/internal/calendarpb"
	"github.com/alexdemen/go-otus/calendar/internal/core"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EventServer struct {
	store core.Store
}

func NewEventServer(store core.Store) EventServer {
	return EventServer{
		store: store,
	}
}

func (e EventServer) Save(cxt context.Context, event *calendarpb.Event) (*calendarpb.Event, error) {
	coreEvent, err := convertFromPB(*event)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if coreEvent.Id != 0 {
		err = e.store.Edit(cxt, coreEvent)

	} else {
		coreEvent, err = e.store.Add(cxt, coreEvent)
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	event.Id = coreEvent.Id
	return event, nil
}

func (e EventServer) Remove(cxt context.Context, event *calendarpb.Event) (*calendarpb.Result, error) {
	coreEvent, err := convertFromPB(*event)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = e.store.Remove(cxt, coreEvent)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &calendarpb.Result{Success: true}, nil
}

func (e EventServer) List(cxt context.Context, user *calendarpb.User) (*calendarpb.EventList, error) {
	list, err := e.store.List(cxt)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	res := calendarpb.EventList{
		Events: make([]*calendarpb.Event, 0, len(list)),
	}

	for _, val := range list {
		event, err := convertToPB(val)
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		res.Events = append(res.Events, event)
	}

	return &res, nil
}

func convertToPB(event core.Event) (*calendarpb.Event, error) {
	startDate, err := ptypes.TimestampProto(event.StartDate)
	duration := ptypes.DurationProto(event.Duration)
	if err != nil {
		return nil, err
	}
	return &calendarpb.Event{
		Id:          event.Id,
		Name:        event.Name,
		Description: event.Description,
		StartDate:   startDate,
		Duration:    duration,
	}, nil
}

func convertFromPB(event calendarpb.Event) (core.Event, error) {
	startDate, err := ptypes.Timestamp(event.StartDate)
	if err != nil {
		return core.Event{}, err
	}
	duration, err := ptypes.Duration(event.Duration)
	if err != nil {
		return core.Event{}, err
	}
	return core.Event{
		Id:          event.Id,
		Name:        event.Name,
		Description: event.Description,
		StartDate:   startDate,
		Duration:    duration,
	}, nil
}
