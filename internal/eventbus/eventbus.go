package eventbus

import "github.com/piotrowski/ebitris/internal/event"

type Event struct {
	Type  event.Type
	Value int
}

type Handler func(Event)

type Bus struct {
	handlers map[event.Type][]Handler
	queue    []Event
}

func NewBus() *Bus {
	return &Bus{handlers: make(map[event.Type][]Handler)}
}

func (b *Bus) Subscribe(t event.Type, h Handler) {
	b.handlers[t] = append(b.handlers[t], h)
}

func (b *Bus) Emit(e Event) {
	b.queue = append(b.queue, e)
}

func (b *Bus) Dispatch() {
	for _, e := range b.queue {
		for _, h := range b.handlers[e.Type] {
			h(e)
		}
	}
	b.queue = b.queue[:0]
}
