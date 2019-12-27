package doublyLinkedList

import "testing"

func TestList_First(t *testing.T) {
	list := List{}

	if list.Len() != 0 || list.first != nil || list.last != nil {
		t.Errorf("Ошибка инициализации списка.")
	}

	list.PushFront("first")

	if list.Len() == 0 || list.first == nil || list.last == nil {
		t.Errorf("Ошибка добавления первого элемента.")
	}

	if item, err := list.First(); err == nil {
		if val, ok := item.data.(string); !ok || val != "first" {
			t.Errorf("Значение не равно заданному раннее.")
		}
	} else {
		t.Errorf("Не удалось получить первый элемент.")
	}
}
