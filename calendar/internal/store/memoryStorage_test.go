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

func TestMemoryStore_List(t *testing.T) {
	storage := NewMemoryStore()

	testEvent := core.Event{Name: "test",
		StartDate:  time.Date(2020, 1, 1, 1, 1, 0, 0, time.UTC),
		FinishDate: time.Date(2020, 1, 1, 2, 0, 0, 0, time.UTC)}

	err := storage.Add(&testEvent)
	if err != nil {
		t.Fatal("Error in method Add.")
	}

	list, err := storage.List()
	if err != nil {
		t.Fatal("Error getting event list.")
	}

	if len(list) == 0 {
		t.Fatal("Event list is empty.")
	}

	if val, ok := storage.events[list[0].Id]; ok {
		if val.Name != testEvent.Name {
			t.Fatal("Error getting event.")
		}
	} else {
		t.Fatal("Error getting event.")
	}
}

func TestMemoryStore_ListEmpty(t *testing.T) {
	storage := NewMemoryStore()

	if _, err := storage.List(); err != nil {
		if !errors.Is(err, core.ErrNoEvents) {
			t.Fatal("Error getting event list.")
		}
	} else {
		t.Fatal("Event list is not empty.")
	}
}

func TestMemoryStore_Remove(t *testing.T) {
	storage := NewMemoryStore()

	testEvent := core.Event{Name: "test",
		StartDate:  time.Date(2020, 1, 1, 1, 1, 0, 0, time.UTC),
		FinishDate: time.Date(2020, 1, 1, 2, 0, 0, 0, time.UTC)}

	err := storage.Add(&testEvent)
	if err != nil {
		t.Fatal("Error in method Add.")
	}

	if len(storage.events) == 0 {
		t.Fatal("Event list is empty.")
	}

	if err = storage.Remove(testEvent); err != nil {
		t.Fatal("Error deleting event.")
	}

	if len(storage.events) != 0 {
		t.Fatal("Event list is not empty.")
	}
}

func TestMemoryStore_RemoveInEmptyList(t *testing.T) {
	storage := NewMemoryStore()

	testEvent := core.Event{Name: "test",
		StartDate:  time.Date(2020, 1, 1, 1, 1, 0, 0, time.UTC),
		FinishDate: time.Date(2020, 1, 1, 2, 0, 0, 0, time.UTC)}

	if len(storage.events) != 0 {
		t.Fatal("Event list is not empty.")
	}

	if err := storage.Remove(testEvent); err == nil {
		t.Fatal("Error deleting event in empty list.")
	}
}
