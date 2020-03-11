package postgres

import (
	"context"
	"github.com/alexdemen/go-otus/calendar/internal/core"
	"testing"
	"time"
)

func TestStore_Add(t *testing.T) {
	store, err := NewStore("postgres://event_admin:123@localhost:5432/eventsdb")
	if err != nil {
		t.Fatal(err)
	}
	defer store.database.Close()

	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		t.Fatal(err)
	}
	start := time.Date(2020, 3, 11, 18, 30, 0, 0, loc)

	event := core.Event{
		Name:        "Test1",
		Description: "",
		StartDate:   start,
		Duration:    time.Until(start.Add(time.Hour)),
	}

	event, err = store.Add(context.Background(), event)
	if err != nil {
		t.Fatal(err)
	}
}
