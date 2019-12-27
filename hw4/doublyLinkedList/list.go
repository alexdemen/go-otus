package doublyLinkedList

import (
	"errors"
)

type List struct {
	len   int
	first *Item
	last  *Item
}

func (l List) Len() int {
	return l.len
}

func (l List) First() (*Item, error) {
	if l.len == 0 {
		return nil, errors.New("Список не содержит элементы.")
	}

	return l.first, nil
}

func (l List) Last() (*Item, error) {
	if l.len == 0 {
		return nil, errors.New("Список не содержит элементы.")
	}

	return l.last, nil
}

func (l *List) PushFront(v interface{}) {
	if first, err := l.First(); err != nil {
		item := &Item{data: v}
		l.first = item
		l.last = item
	} else {
		newItem := &Item{data: v, next: first}
		first.prev = newItem
		l.first = newItem
	}
	l.len++
}

func (l *List) PushBack(v interface{}) {
	if prev, err := l.Last(); err != nil {
		item := &Item{data: v}
		l.first = item
		l.last = item
	} else {
		newItem := &Item{data: v, prev: prev}
		prev.next = newItem
		l.last = newItem
	}
	l.len++
}

func (l *List) Remove(item Item) {
	prev, next := item.prev, item.next

	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	l.len--
}
