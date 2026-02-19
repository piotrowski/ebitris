package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/input"
	"github.com/piotrowski/ebitris/internal/render"
)

type PauseScene struct {
	manager *Manager
	input   *input.InputManager
}

func NewPauseScene(manager *Manager) *PauseScene {
	return &PauseScene{
		manager: manager,
		input:   input.NewInputManager(),
	}
}

func (s *PauseScene) Update() error {
	if s.input.IsKeyJustPressed(ebiten.KeyEscape) {
		s.manager.SwitchBack()
	}
	return nil
}

func (s *PauseScene) Draw(screen *ebiten.Image) {
	font := render.GetDefaultFont(render.FontLarge)

	render.DrawText(screen, "PAUSED", 260, 400, font)
}

func (s *PauseScene) OnEnter() {}
func (s *PauseScene) OnExit()  {}
