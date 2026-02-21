package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/input"
	"github.com/piotrowski/ebitris/internal/render"
	"github.com/piotrowski/ebitris/internal/ui"
)

type MainMenuScene struct {
	manager *Manager
	input   *input.InputManager
	menu    *ui.Menu
}

func NewMainMenuScene(manager *Manager) *MainMenuScene {
	return &MainMenuScene{
		manager: manager,
		input:   input.NewInputManager(),
		menu:    ui.NewMenu([]string{"Start Game", "Scoreboard", "Exit"}),
	}
}

func (s *MainMenuScene) Update() error {
	if s.menu.HandleInput(s.input) {
		switch s.menu.Selected() {
		case 0:
			s.manager.SwitchTo(NewStandardGameplayScene(s.manager))
		case 1:
			// s.manager.SwitchTo(NewScoreboardScene(s.manager))
		case 2:
			s.manager.Quit()
		}
	}

	return nil
}

func (s *MainMenuScene) Draw(screen *ebiten.Image) {
	fontLarge := render.GetDefaultFont(render.FontLarge)
	render.DrawText(screen, "Ebitris", 5, 5, fontLarge)
	s.menu.Draw(screen, 5, 10)
}

func (s *MainMenuScene) OnEnter() {}
func (s *MainMenuScene) OnExit()  {}
