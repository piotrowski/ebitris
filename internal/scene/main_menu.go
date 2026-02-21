package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/input"
	"github.com/piotrowski/ebitris/internal/render"
)

type MainMenuScene struct {
	manager *Manager
	input   *input.InputManager

	focus int // 0 = Start Game, 1 = Scoreboard, 2 = Exit
}

func NewMainMenuScene(manager *Manager) *MainMenuScene {
	return &MainMenuScene{
		manager: manager,
		input:   input.NewInputManager(),
	}
}

func (s *MainMenuScene) Update() error {
	if s.input.IsKeyJustPressed(ebiten.KeyDown) {
		s.focus = (s.focus + 1) % 3
	}
	if s.input.IsKeyJustPressed(ebiten.KeyUp) {
		s.focus = (s.focus - 1 + 3) % 3
	}
	if s.input.IsKeyJustPressed(ebiten.KeyEnter) {
		switch s.focus {
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
	fontMedium := render.GetDefaultFont(render.FontMedium)

	render.DrawText(screen, "Ebitris", 5, 5, fontLarge)
	render.DrawRectangle(screen, 4, 10+s.focus, 5, 1, color.RGBA{16, 16, 16, 255})

	render.DrawText(screen, "Start Game", 5, 10, fontMedium)
	render.DrawText(screen, "Scoreboard", 5, 11, fontMedium)
	render.DrawText(screen, "Exit", 5, 12, fontMedium)
}

func (s *MainMenuScene) OnEnter() {}
func (s *MainMenuScene) OnExit()  {}
