package event

type Type string

const (
	EventTypeLineClear Type = "line_clear"
	EventTypeLevelUp   Type = "level_up"
	EventTypeGameOver  Type = "game_over"
)
