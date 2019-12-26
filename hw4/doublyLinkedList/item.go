package doublyLinkedList

type Item struct {
	data interface{}
	next *Item
	prev *Item
}

func (i Item) Value() interface{} {
	return i.data
}

func (i Item) Next() *Item {
	return i.next
}

func (i Item) Prev() *Item {
	return i.prev
}
