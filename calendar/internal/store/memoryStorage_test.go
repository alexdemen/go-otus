package store

import (
	"errors"
	"github.com/alexdemen/go-otus/calendar/internal/core"
	"testing"
	"time"
)

func TestMemoryStore_Add(t *testing.T) {
	storage := NewMemoryStore()

	testEvent := core.Event{Name: "test",
		StartDate:  time.Date(2020, 1, 1, 1, 1, 0, 0, time.UTC),
		FinishDate: time.Date(2020, 1, 1, 2, 0, 0, 0, time.UTC)}

	err := storage.Add(&testEvent)
	if err != nil {
		t.Fatal("Error in method Add.")
	}

	if testEvent.Id != 1 {
		t.Fatal("Index assignment error.")
	}
}

func TestMemoryStore_AddWithoutIntersection(t *testing.T) {
	storage := NewMemoryStore()

	firstEvent := core.Event{Name: "test",
		StartDate:  time.Date(2020, 1, 1, 1, 1, 0, 0, time.UTC),
		FinishDate: time.Date(2020, 1, 1, 2, 0, 0, 0, time.UTC)}

	secondEvent := core.Event{Name: "test",
		StartDate:  time.Date(2020, 1, 1, 2, 1, 0, 0, time.UTC),
		FinishDate: time.Date(2020, 1, 1, 2, 30, 0, 0, time.UTC)}

	err := storage.Add(&firstEvent)
	if err != nil {
		t.Fatal("Error in method Add.")
	}

	err = storage.Add(&secondEvent)
	if err != nil {
		t.Fatal("Error in method Add.")
	}

	if len(storage.events) < 2 {
		t.Fatal("Invalid len of events list.")
	}
}

func TestMemoryStore_AddWithIntersection(t *testing.T) {
	storage := NewMemoryStore()

	firstEvent := core.Event{Name: "test",
		StartDate:  time.Date(2020, 1, 1, 1, 1, 0, 0, time.UTC),
		FinishDate: time.Date(2020, 1, 1, 2, 0, 0, 0, time.UTC)}

	secondEvent := core.Event{Name: "test",
		StartDate:  time.Date(2020, 1, 1, 1, 30, 0, 0, time.UTC),
		FinishDate: time.Date(2020, 1, 1, 2, 30, 0, 0, time.UTC)}

	err := storage.Add(&firstEvent)
	if err != nil {
		t.Fatal("Error in method Add.")
	}

	err = storage.Add(&secondEvent)
	if err != nil {
		if !errors.Is(err, core.ErrDateBusy) {
			t.Fatal("Error in method Add, when time interval intersection.")
		}
	} else {
		t.Fatal("Incorrect intersection detection.")
	}

	if len(storage.events) != 1 {
		t.Fatal("Invalid len of events list.")
	}
}
