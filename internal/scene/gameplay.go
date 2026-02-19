package scene

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/input"
	"github.com/piotrowski/ebitris/internal/render"
	"github.com/piotrowski/ebitris/internal/tetris"
)

type GameplayScene struct {
	manager *Manager
	state   *tetris.GameState
	input   *input.InputManager
}

func NewGameplayScene(manager *Manager, width, height int) *GameplayScene {
	return &GameplayScene{
		manager: manager,
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
		s.manager.SwitchTo(NewPauseScene(s.manager))
		return nil
	}

	// Update game state (gravity, etc.)
	s.state.Update()

	return nil
}

func (s *GameplayScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 10, G: 10, B: 20, A: 255})

	offsetX, offsetY := 5, 2
	render.DrawBoard(screen, s.state.GetBoard(), offsetX, offsetY)
	render.DrawPiece(screen, s.state.GetCurrentPiece(), offsetX, offsetY)

	font := render.GetDefaultFont(render.FontMedium)

	render.DrawText(screen, fmt.Sprintf("Score: %d", s.state.GetScore()), 10, 30, font)
	render.DrawText(screen, fmt.Sprintf("Level: %d", s.state.GetLevel()), 10, 45, font)
	render.DrawText(screen, fmt.Sprintf("Lines: %d", s.state.GetLinesCleared()), 10, 60, font)

	render.DrawText(screen, "Next:", 470, 200, font)
	render.DrawPiece(screen, s.state.GetNextPiece(), 13, 8)
}

func (s *GameplayScene) OnEnter() {
	s.state.Resume()
}

func (s *GameplayScene) OnExit() {
	s.state.Pause()
}
