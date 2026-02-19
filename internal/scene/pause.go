package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/input"
	"github.com/piotrowski/ebitris/internal/render"
	"github.com/piotrowski/ebitris/internal/tetris"
)

type PauseScene struct {
	manager  *Manager
	gameplay *GameplayScene
	input    *input.InputManager
}

func NewPauseScene(manager *Manager, gameplay *GameplayScene) *PauseScene {
	return &PauseScene{
		manager:  manager,
		gameplay: gameplay,
		input:    input.NewInputManager(),
	}
}

func (s *PauseScene) Update() error {
	if s.input.IsKeyJustPressed(ebiten.KeyEscape) {
		s.gameplay.state.Status = tetris.StatusPlaying
		s.manager.SwitchTo(s.gameplay)
	}
	return nil
}

func (s *PauseScene) Draw(screen *ebiten.Image) {
	s.gameplay.Draw(screen)

	font := render.GetDefaultFont(render.FontLarge)

	render.DrawText(screen, "PAUSED", 260, 400, font)
}

func (s *PauseScene) OnEnter() {}
func (s *PauseScene) OnExit()  {}
