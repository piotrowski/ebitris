package scene

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/input"
	"github.com/piotrowski/ebitris/internal/render"
)

type GameOverScene struct {
	manager *Manager
	input   *input.InputManager
	score   int
}

func NewGameOverScene(manager *Manager, score int) *GameOverScene {
	return &GameOverScene{
		manager: manager,
		score:   score,
		input:   input.NewInputManager(),
	}
}

func (s *GameOverScene) Update() error {
	if s.input.IsKeyJustPressed(ebiten.KeyR) {
		s.manager.SwitchTo(NewMainMenuScene(s.manager))
	}
	return nil
}

func (s *GameOverScene) Draw(screen *ebiten.Image) {
	font := render.GetDefaultFont(render.FontLarge)

	render.DrawText(screen, "Game Over", 5, 5, font)
	render.DrawText(screen, fmt.Sprintf("Score: %d", s.score), 5, 6, font)
}

func (s *GameOverScene) OnEnter() {}
func (s *GameOverScene) OnExit()  {}
