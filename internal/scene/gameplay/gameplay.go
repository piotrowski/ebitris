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
	if s.input.ShouldMove(ebiten.KeyLeft) {
		s.state.MoveLeft()
	}
	if s.input.ShouldMove(ebiten.KeyRight) {
		s.state.MoveRight()
	}
	if s.input.IsKeyJustPressed(ebiten.KeyUp) {
		s.state.Rotate()
	}
	if s.input.ShouldMove(ebiten.KeyDown) {
		s.state.MoveDown()
	}
	if s.input.IsKeyJustPressed(ebiten.KeySpace) {
		s.state.HardDrop()
	}
	if s.input.IsKeyJustPressed(ebiten.KeyEscape) {
		s.emitter.Emit(event.Event{Type: event.EventTypePause})
		return nil
	}

	if s.state.IsGameOver() {
		s.emitter.Emit(event.Event{Type: event.EventTypeGameOver, Payload: map[string]int{"score": 123}})
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
	// s.manager.audioManager.SetPlaylist(music.ReturnOfThe8BitEra, music.ArcadeBeat)
	// go s.manager.audioManager.StartAutoPlay()
	s.state.Resume()
}

func (s *GameplayScene) OnExit() {
	// go s.manager.audioManager.StopAutoPlay()

	s.state.Pause()
}
