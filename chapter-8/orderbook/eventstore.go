package orderbook

// === Creating the EventStore Interface === //
type EventStore interface {
	Append(Event)
	Load() []Event
}

// === Implementing Memory EventStore === //
type MemoryEventStore struct {
	events []Event
}

func NewMemoryEventStore() *MemoryEventStore {
	return &MemoryEventStore{
		events: make([]Event, 0),
	}
}

func (s *MemoryEventStore) Append(event Event) {
	s.events = append(s.events, event)
}

func (s *MemoryEventStore) Load() []Event {
	return s.events
}
