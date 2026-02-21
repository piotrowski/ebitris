package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/input"
	"github.com/piotrowski/ebitris/internal/render"
	"github.com/piotrowski/ebitris/internal/ui"
)

type PauseScene struct {
	manager *Manager
	input   *input.InputManager
	menu    *ui.Menu
}

func NewPauseScene(manager *Manager) *PauseScene {
	return &PauseScene{
		manager: manager,
		input:   input.NewInputManager(),
		menu:    ui.NewMenu([]string{"Resume", "Restart", "Main Menu"}),
	}
}

func (s *PauseScene) Update() error {
	if s.input.IsKeyJustPressed(ebiten.KeyEscape) {
		s.manager.SwitchBack()
	}

	if s.menu.HandleInput(s.input) {
		switch s.menu.Selected() {
		case 0:
			s.manager.SwitchBack()
		case 1:
			s.manager.SwitchTo(NewStandardGameplayScene(s.manager))
		case 2:
			s.manager.SwitchTo(NewMainMenuScene(s.manager))
		}
	}
	return nil
}

func (s *PauseScene) Draw(screen *ebiten.Image) {
	fontLarge := render.GetDefaultFont(render.FontLarge)
	render.DrawText(screen, "PAUSED", 5, 5, fontLarge)
	s.menu.Draw(screen, 5, 10)
}

func (s *PauseScene) OnEnter() {}
func (s *PauseScene) OnExit()  {}
