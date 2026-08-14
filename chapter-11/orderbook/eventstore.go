package orderbook

// === Creating the EventStore Interface === //
type EventStore interface {
	Append(Event)
	Load() []Event
	LoadAfter(lastEventID int) []Event
	LastEventID() int
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

func (s *MemoryEventStore) LastEventID() int {
	return len(s.events)
}

func (s *MemoryEventStore) LoadAfter(lastEventID int) []Event {
	if lastEventID >= len(s.events) {
		return nil
	}
	return s.events[lastEventID:]
}
