package store

import (
	"github.com/alexdemen/go-otus/calendar/internal/core"
	"time"
)

type memoryStore struct {
	events map[int64]core.Event
	idSeq  int64
}

func NewMemoryStore() memoryStore {
	return memoryStore{
		events: make(map[int64]core.Event),
		idSeq:  1,
	}
}

func (m *memoryStore) Add(event *core.Event) error {
	if exist := m.existIntersection(*event); exist {
		return core.ErrDateBusy
	}
	event.Id = m.idSeq
	m.idSeq++
	m.events[event.Id] = *event
	return nil
}

func (m *memoryStore) Edit(event core.Event) error {
	if exist := m.existIntersection(event); exist {
		return core.ErrDateBusy
	}
	if _, ok := m.events[event.Id]; !ok {
		return core.ErrEventNotExist{}
	}

	m.events[event.Id] = event
	return nil
}

func (m *memoryStore) Remove(event core.Event) error {
	if _, ok := m.events[event.Id]; !ok {
		return core.ErrEventNotExist{}
	}
	delete(m.events, event.Id)
	return nil
}

func (m memoryStore) List() ([]core.Event, error) {
	if len(m.events) == 0 {
		return nil, core.ErrNoEvents{}
	}
	res := make([]core.Event, 0, len(m.events))
	for _, val := range m.events {
		res = append(res, val)
	}
	return res, nil
}

func (m memoryStore) existIntersection(event core.Event) bool {
	inInterval := func(target time.Time, start time.Time, end time.Time) bool {
		return (target.After(start) && target.Before(end)) || target.Equal(start) || target.Equal(end)
	}

	for key, val := range m.events {
		if (inInterval(event.StartDate, val.StartDate, val.FinishDate) ||
			inInterval(event.FinishDate, val.StartDate, val.FinishDate)) && event.Id != key {
			return true
		}
	}
	return false
}
