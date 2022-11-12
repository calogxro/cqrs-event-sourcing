package es

type EventStore struct {
	store map[string][]IEvent
}

func NewEventStore() *EventStore {
	return &EventStore{
		store: make(map[string][]IEvent),
	}
}

func (es *EventStore) getEvents(userId string) []IEvent {
	return es.store[userId]
}

func (es *EventStore) AddEvent(userId string, event IEvent) {
	userEvents := es.store[userId]
	es.store[userId] = append(userEvents, event)
}
