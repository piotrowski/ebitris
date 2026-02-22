package eventbus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscribeAndDispatch(t *testing.T) {
	t.Parallel()

	bus := NewBus()

	var received []Event
	bus.Subscribe("test", func(e Event) {
		received = append(received, e)
	})

	bus.Emit(Event{Type: "test", Value: 42})
	bus.Dispatch()

	assert.Equal(t, []Event{{Type: "test", Value: 42}}, received)
}

func TestDispatchClearsQueue(t *testing.T) {
	t.Parallel()

	bus := NewBus()

	count := 0
	bus.Subscribe("tick", func(e Event) { count++ })

	bus.Emit(Event{Type: "tick"})
	bus.Dispatch()
	bus.Dispatch() // second dispatch should fire nothing

	assert.Equal(t, 1, count)
}

func TestMultipleHandlersSameType(t *testing.T) {
	t.Parallel()

	bus := NewBus()

	calls := make([]int, 2)
	bus.Subscribe("ev", func(e Event) { calls[0]++ })
	bus.Subscribe("ev", func(e Event) { calls[1]++ })

	bus.Emit(Event{Type: "ev"})
	bus.Dispatch()

	assert.Equal(t, []int{1, 1}, calls)
}

func TestHandlersOnlyReceiveMatchingType(t *testing.T) {
	t.Parallel()

	bus := NewBus()

	var aEvents, bEvents []Event
	bus.Subscribe("a", func(e Event) { aEvents = append(aEvents, e) })
	bus.Subscribe("b", func(e Event) { bEvents = append(bEvents, e) })

	bus.Emit(Event{Type: "a", Value: 1})
	bus.Emit(Event{Type: "b", Value: 2})
	bus.Dispatch()

	assert.Equal(t, []Event{{Type: "a", Value: 1}}, aEvents)
	assert.Equal(t, []Event{{Type: "b", Value: 2}}, bEvents)
}

func TestEmitMultipleBeforeDispatch(t *testing.T) {
	t.Parallel()

	bus := NewBus()

	var values []int
	bus.Subscribe("score", func(e Event) { values = append(values, e.Value) })

	bus.Emit(Event{Type: "score", Value: 10})
	bus.Emit(Event{Type: "score", Value: 20})
	bus.Emit(Event{Type: "score", Value: 30})
	bus.Dispatch()

	assert.Equal(t, []int{10, 20, 30}, values)
}

func TestNoHandlerForType(t *testing.T) {
	t.Parallel()

	bus := NewBus()
	// Emitting with no subscriber should not panic
	bus.Emit(Event{Type: "unhandled", Value: 99})
	assert.NotPanics(t, func() { bus.Dispatch() })
}

func TestQueueIsEmptyAfterDispatch(t *testing.T) {
	bus := NewBus()

	bus.Emit(Event{Type: "x"})
	bus.Dispatch()

	assert.Empty(t, bus.queue)
}
