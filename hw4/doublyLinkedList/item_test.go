package doublyLinkedList

import "testing"

func TestItem_Value(t *testing.T) {
	item := Item{data: "ok"}

	if val, ok := item.Value().(string); !ok || val != "ok" {
		t.Fatalf("Ошибка получения значения элемента списка.")
	}
}

func TestItem_Next(t *testing.T) {
	first := Item{data: "first", next: &Item{data: "last"}}

	next := first.Next()

	if str := next.Value().(string); str != "last" {
		t.Fatalf("Ошибка получение значения следующего элемента.")
	}
}

func TestItem_Prev(t *testing.T) {
	last := Item{data: "last", prev: &Item{data: "first"}}

	prev := last.Prev()

	if str := prev.Value().(string); str != "first" {
		t.Fatalf("Ошибка получение значения предыдущего элемента.")
	}
}
