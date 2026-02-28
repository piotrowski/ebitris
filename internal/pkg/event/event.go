package event

type EventType int

const (
	EventTypeLineClear EventType = iota
	EventTypeLevelUp
	EventTypeGameOver

	EventTypeGoBack
	EventTypeStartGame
	EventTypeMainMenu
	EventTypeScoreboard

	EventTypePause
	EventTypeQuit

	EventTypeBlockPlaced
	EventTypeBlockMovedByPlayer
)

type Emitter interface {
	Emit(e Event)
}

type Subscriber interface {
	Subscribe(t EventType, h Handler)
}

type EmitterSubscriber interface {
	Emitter
	Subscriber
}

type Dispatcher interface {
	Dispatch()
}

type GameOverPayload struct {
	Score int
	Lines int
	Level int
}
