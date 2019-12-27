package doublyLinkedList

import "testing"

func TestList_First(t *testing.T) {
	list := List{}

	if list.Len() != 0 || list.first != nil || list.last != nil {
		t.Fatalf("Ошибка инициализации списка.")
	}

	list.PushFront("first")

	if list.Len() == 0 || list.first == nil || list.last == nil {
		t.Fatalf("Ошибка добавления первого элемента.")
	}

	if item, err := list.First(); err == nil {
		if val, ok := item.data.(string); !ok || val != "first" {
			t.Fatalf("Значение не равно заданному раннее.")
		}
	} else {
		t.Fatalf("Не удалось получить первый элемент.")
	}
}

func TestList_Last(t *testing.T) {
	list := List{}

	if list.Len() != 0 || list.first != nil || list.last != nil {
		t.Fatalf("Ошибка инициализации списка.")
	}

	list.PushBack("back")

	if list.Len() == 0 || list.first == nil || list.last == nil {
		t.Fatalf("Ошибка добавления первого элемента.")
	}

	if item, err := list.Last(); err == nil {
		if val, ok := item.data.(string); !ok || val != "back" {
			t.Fatalf("Значение не равно заданному раннее.")
		}
	} else {
		t.Fatalf("Не удалось получить первый элемент.")
	}
}

func TestThreeValue(t *testing.T) {
	list := List{}

	if list.Len() != 0 || list.first != nil || list.last != nil {
		t.Fatalf("Ошибка инициализации списка.")
	}

	list.PushFront("first")
	list.PushFront("second")

	if val, err := list.First(); err != nil {
		t.Fatalf("Не удалось вставить второй элемент.")
	} else if str, ok := val.data.(string); !ok || str != "second" {
		t.Fatalf("Ошибка вставки элемента.")
	}

	list.PushBack("last")

	if val, err := list.Last(); err != nil {
		t.Fatalf("Не удалось вставить второй элемент.")
	} else if str, ok := val.data.(string); !ok || str != "last" {
		t.Fatalf("Ошибка вставки элемента.")
	}

	if list.Len() != 3 {
		t.Fatalf("Неверная длина списка.")
	}
}

func TestList_Remove(t *testing.T) {
	list := List{}

	if list.Len() != 0 || list.first != nil || list.last != nil {
		t.Fatalf("Ошибка инициализации списка.")
	}

	list.PushFront("first")
	list.PushFront("second")
	list.PushBack("last")

	remItem, _ := list.First()
	remItem = remItem.Next()

	list.Remove(*remItem)

	if list.Len() != 2 {
		t.Fatalf("Неверная длина списка.")
	}

	first, _ := list.First()
	if str, ok := first.data.(string); !ok || str != "second" {
		t.Fatalf("Первый элемент отличается от ожидаемого.")
	}

	last, _ := list.Last()
	if str, ok := last.data.(string); !ok || str != "last" {
		t.Fatalf("Последний элемент отличается от ожидаемого.")
	}
}
