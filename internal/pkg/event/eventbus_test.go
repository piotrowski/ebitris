package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscribeAndDispatch(t *testing.T) {
	t.Parallel()

	bus := NewEventManager()

	var received []Event
	bus.Subscribe(EventTypeGameOver, func(e Event) {
		received = append(received, e)
	})

	bus.Emit(Event{Type: EventTypeGameOver, Payload: 42})
	bus.Dispatch()

	assert.Equal(t, []Event{{Type: EventTypeGameOver, Payload: 42}}, received)
}

func TestDispatchClearsQueue(t *testing.T) {
	t.Parallel()

	bus := NewEventManager()

	count := 0
	bus.Subscribe(EventTypeGameOver, func(e Event) { count++ })

	bus.Emit(Event{Type: EventTypeGameOver})
	bus.Dispatch()
	bus.Dispatch() // second dispatch should fire nothing

	assert.Equal(t, 1, count)
}

func TestMultipleHandlersSameType(t *testing.T) {
	t.Parallel()

	bus := NewEventManager()

	calls := make([]int, 2)
	bus.Subscribe(EventTypeGameOver, func(e Event) { calls[0]++ })
	bus.Subscribe(EventTypeGameOver, func(e Event) { calls[1]++ })

	bus.Emit(Event{Type: EventTypeGameOver})
	bus.Dispatch()

	assert.Equal(t, []int{1, 1}, calls)
}

func TestHandlersOnlyReceiveMatchingType(t *testing.T) {
	t.Parallel()

	bus := NewEventManager()

	var aEvents, bEvents []Event
	bus.Subscribe(EventTypeGameOver, func(e Event) { aEvents = append(aEvents, e) })
	bus.Subscribe(EventTypeGoBack, func(e Event) { bEvents = append(bEvents, e) })

	bus.Emit(Event{Type: EventTypeGameOver, Payload: 1})
	bus.Emit(Event{Type: EventTypeGoBack, Payload: 2})
	bus.Dispatch()

	assert.Equal(t, []Event{{Type: EventTypeGameOver, Payload: 1}}, aEvents)
	assert.Equal(t, []Event{{Type: EventTypeGoBack, Payload: 2}}, bEvents)
}

func TestEmitMultipleBeforeDispatch(t *testing.T) {
	t.Parallel()

	bus := NewEventManager()

	var values []int
	bus.Subscribe(EventTypeGameOver, func(e Event) { values = append(values, e.Payload.(int)) })

	bus.Emit(Event{Type: EventTypeGameOver, Payload: 10})
	bus.Emit(Event{Type: EventTypeGameOver, Payload: 20})
	bus.Emit(Event{Type: EventTypeGameOver, Payload: 30})
	bus.Dispatch()

	assert.Equal(t, []int{10, 20, 30}, values)
}

func TestNoHandlerForType(t *testing.T) {
	t.Parallel()

	bus := NewEventManager()
	// Emitting with no subscriber should not panic
	bus.Emit(Event{Type: EventTypeGameOver, Payload: 99})
	assert.NotPanics(t, func() { bus.Dispatch() })
}

func TestQueueIsEmptyAfterDispatch(t *testing.T) {
	bus := NewEventManager()

	bus.Emit(Event{Type: EventTypeGameOver})
	bus.Dispatch()

	assert.Empty(t, bus.queue)
}
