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

	EventTypeStartPlaylist
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
