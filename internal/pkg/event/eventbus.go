package event

type Event struct {
	Type    EventType
	Payload any
}

type Handler func(Event)

type EventManager struct {
	handlers map[EventType][]Handler
	queue    []Event
}

func NewEventManager() *EventManager {
	return &EventManager{handlers: make(map[EventType][]Handler)}
}

func (b *EventManager) Subscribe(t EventType, h Handler) {
	b.handlers[t] = append(b.handlers[t], h)
}

func (b *EventManager) Emit(e Event) {
	b.queue = append(b.queue, e)
}

func (b *EventManager) Dispatch() {
	for _, e := range b.queue {
		for _, h := range b.handlers[e.Type] {
			h(e)
		}
	}
	b.queue = b.queue[:0]
}
