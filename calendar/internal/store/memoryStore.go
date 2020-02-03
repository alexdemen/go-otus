package store

type memoryStore struct {
	events map[int64]Event
}

func NewMemoryStore() memoryStore {
	return memoryStore{
		events: make(map[int64]Event),
	}
}

func (m *memoryStore) Add(event Event) error {
	//TODO
	m.events[event.Id] = event
	return nil
}

func (m *memoryStore) Edit(event Event) error {
	//TODO
	m.events[event.Id] = event
	return nil
}

func (m *memoryStore) Remove(event Event) error {
	//TODO
	delete(m.events, event.Id)
	return nil
}

func (m memoryStore) List() ([]Event, error) {
	res := make([]Event, 0, len(m.events))
	//TODO
	return res, nil
}
