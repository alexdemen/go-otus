package postgres

import (
	"context"
	"fmt"
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

	//test := "tea"
	event := core.Event{
		Name:        "Test1",
		Description: nil, //&test,
		StartDate:   start,
		Duration:    start.Add(time.Hour).Sub(start), //Sub(start.Add(time.Hour)),
	}

	event, err = store.Add(context.Background(), event)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStore_List(t *testing.T) {
	store, err := NewStore("postgres://event_admin:123@localhost:5432/eventsdb")
	if err != nil {
		t.Fatal(err)
	}
	defer store.database.Close()

	data, _ := store.List(context.Background())

	for _, ev := range data {
		endTime := ev.StartDate.Add(ev.Duration)
		fmt.Print(endTime)
	}
}

func TestStore_Remove(t *testing.T) {
	store, err := NewStore("postgres://event_admin:123@localhost:5432/eventsdb")
	if err != nil {
		t.Fatal(err)
	}
	defer store.database.Close()

	event := core.Event{
		Id:          10,
		Name:        "Test1",
		Description: nil, //&test,
	}

	store.Remove(context.Background(), event)
}

func TestStore_Edit(t *testing.T) {
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

	test := "tea"
	event := core.Event{
		Id:          10,
		Name:        "Test2",
		Description: &test,
		StartDate:   start,
		Duration:    start.Add(time.Hour).Sub(start), //Sub(start.Add(time.Hour)),
	}

	store.Edit(context.Background(), event)
}
