package gameplay

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/pkg/event"
	"github.com/piotrowski/ebitris/internal/pkg/input"
	"github.com/piotrowski/ebitris/internal/render"
	"github.com/piotrowski/ebitris/internal/tetris"
)

type GameplayScene struct {
	emitter event.Emitter
	state   *tetris.GameState
	input   *input.InputManager
}

func NewStandardGameplayScene(emitter event.Emitter) *GameplayScene {
	return NewGameplayScene(emitter, 10, 20)
}

func NewGameplayScene(emitter event.Emitter, width, height int) *GameplayScene {
	return &GameplayScene{
		emitter: emitter,
		state:   tetris.NewGameState(width, height),
		input:   input.NewInputManager(),
	}
}

func (s *GameplayScene) Update() error {
	var blockedMoved bool
	if s.input.ShouldMove(ebiten.KeyLeft) {
		blockedMoved = s.state.MoveLeft()
	}
	if s.input.ShouldMove(ebiten.KeyRight) {
		blockedMoved = s.state.MoveRight()
	}
	if s.input.IsKeyJustPressed(ebiten.KeyUp) {
		blockedMoved = s.state.Rotate()
	}
	if s.input.ShouldMove(ebiten.KeyDown) {
		blockedMoved = s.state.MoveDown()
	}

	if blockedMoved {
		s.emitter.Emit(event.Event{Type: event.EventTypeBlockMovedByPlayer})
	}

	if s.input.IsKeyJustPressed(ebiten.KeySpace) {
		s.state.HardDrop()
		s.emitter.Emit(event.Event{Type: event.EventTypeBlockPlaced})
	}
	if s.input.IsKeyJustPressed(ebiten.KeyEscape) {
		s.emitter.Emit(event.Event{Type: event.EventTypePause})
		return nil
	}

	if s.state.IsGameOver() {
		s.emitter.Emit(event.Event{Type: event.EventTypeGameOver, Payload: event.GameOverPayload{
			Score: s.state.GetScore(),
			Lines: s.state.GetLinesCleared(),
			Level: s.state.GetLevel(),
		}})
		return nil
	}

	s.state.Update()

	return nil
}

func (s *GameplayScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 10, G: 10, B: 20, A: 255})

	offsetX, offsetY := 4, 2
	render.DrawBoard(screen, s.state.GetBoard(), offsetX, offsetY)
	render.DrawPiece(screen, s.state.GetShadowPiece(), offsetX, offsetY)
	render.DrawPiece(screen, s.state.GetCurrentPiece(), offsetX, offsetY)

	font := render.GetDefaultFont(render.FontMedium)

	render.DrawText(screen, fmt.Sprintf("Score: %d", s.state.GetScore()), 1, 3, font)
	render.DrawText(screen, fmt.Sprintf("Level: %d", s.state.GetLevel()), 1, 4, font)
	render.DrawText(screen, fmt.Sprintf("Lines: %d", s.state.GetLinesCleared()), 1, 5, font)

	render.DrawText(screen, "Next:", 16, 7, font)
	render.DrawPiece(screen, s.state.GetNextPiece(), 12, 9)
}

func (s *GameplayScene) OnEnter() {
	s.state.Resume()
}

func (s *GameplayScene) OnExit() {
	s.state.Pause()
}
