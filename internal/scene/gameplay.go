package scene

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/input"
	"github.com/piotrowski/ebitris/internal/render"
	"github.com/piotrowski/ebitris/internal/tetris"
	"golang.org/x/image/font/basicfont"
)

var font = basicfont.Face7x13

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
		s.tryMove(-1, 0)
	}
	if s.input.ShouldMove(ebiten.KeyRight) {
		s.tryMove(1, 0)
	}
	if s.input.ShouldMove(ebiten.KeyUp) {
		s.tryRotate()
	}
	if s.input.ShouldMove(ebiten.KeyDown) {
		s.tryMove(0, 1)
	}
	if s.input.IsKeyJustPressed(ebiten.KeySpace) {
		s.hardDrop()
	}
	if s.input.IsKeyJustPressed(ebiten.KeyEscape) {
		s.state.Status = tetris.StatusPaused
		s.manager.SwitchTo(NewPauseScene(s.manager, s))
		return nil
	}

	// Update game state (gravity, etc.)
	s.state.Update()

	return nil
}

func (s *GameplayScene) tryMove(dx, dy int) {
	piece := s.state.CurrentPiece
	if !s.state.Board.IsColliding(piece, dx, dy) {
		switch {
		case dx < 0:
			piece.MoveLeft()
		case dx > 0:
			piece.MoveRight()
		}
		if dy > 0 {
			piece.MoveDown()
		}
	}
}

func (s *GameplayScene) tryRotate() {
	piece := s.state.CurrentPiece
	oldRotation := piece.Rotation
	piece.Rotate()

	// Check collision - if invalid, revert
	if s.state.Board.IsColliding(piece, 0, 0) {
		piece.Rotation = oldRotation
	}
}

func (s *GameplayScene) hardDrop() {
	piece := s.state.CurrentPiece
	for !s.state.Board.IsColliding(piece, 0, 1) {
		piece.MoveDown()
	}
}

func (s *GameplayScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 10, G: 10, B: 20, A: 255})

	// Draw board at offset
	offsetX, offsetY := 5, 2
	render.DrawBoard(screen, s.state.Board, offsetX, offsetY)
	render.DrawPiece(screen, s.state.CurrentPiece, offsetX, offsetY)

	// Draw UI
	render.DrawText(screen, fmt.Sprintf("Score: %d", s.state.Score), 10, 30, font)
	render.DrawText(screen, fmt.Sprintf("Level: %d", s.state.Level()), 10, 45, font)
	render.DrawText(screen, fmt.Sprintf("Lines: %d", s.state.LinesCleared), 10, 60, font)

	// Draw next piece preview
	render.DrawText(screen, "Next:", 470, 200, font)
	render.DrawPiece(screen, s.state.NextPiece, 13, 8)
}

func (s *GameplayScene) OnEnter() {
	// Initialize or resume
}

func (s *GameplayScene) OnExit() {
	// Cleanup if needed
}
