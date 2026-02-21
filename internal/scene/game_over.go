package scene

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/input"
	"github.com/piotrowski/ebitris/internal/render"
	"github.com/piotrowski/ebitris/internal/ui"
)

type GameOverScene struct {
	manager *Manager
	input   *input.InputManager
	score   int
	level   int
	lines   int

	menu *ui.Menu

	isInitialsModeActive bool
	initials             string

	scoreSaver ScoreManager
}

func NewGameOverScene(manager *Manager, score, level, lines int) *GameOverScene {
	return &GameOverScene{
		manager: manager,
		score:   score,
		level:   level,
		lines:   lines,
		input:   input.NewInputManager(),

		menu: ui.NewMenu([]string{"Save Score", "Restart", "Main Menu"}),

		scoreSaver: manager.scoreManager,
	}
}

func (s *GameOverScene) Update() error {
	if s.isInitialsModeActive {
		return s.initialsMode()
	}

	if s.menu.HandleInput(s.input) {
		switch s.menu.Selected() {
		case 0:
			s.isInitialsModeActive = true
		case 1:
			s.manager.SwitchTo(NewStandardGameplayScene(s.manager))
		case 2:
			s.manager.SwitchTo(NewMainMenuScene(s.manager))
		}
	}

	return nil
}

func (s *GameOverScene) initialsMode() error {
	if s.input.IsKeyJustPressed(ebiten.KeyEnter) {
		s.scoreSaver.SaveScore(s.initials, s.score, s.level, s.lines)
		s.manager.SwitchTo(NewMainMenuScene(s.manager))
		return nil
	}

	if s.input.IsKeyJustPressed(ebiten.KeyEscape) {
		s.isInitialsModeActive = false
		s.initials = ""
		return nil
	}

	if s.input.IsKeyJustPressed(ebiten.KeyBackspace) && s.initials != "" {
		s.initials = s.initials[:len(s.initials)-1]
	} else {
		for _, key := range s.input.GetJustPressedKeys() {
			if key >= ebiten.KeyA && key <= ebiten.KeyZ && len(s.initials) < 3 {
				s.initials += string('A' + (key - ebiten.KeyA))
			}
		}
	}

	return nil
}

func (s *GameOverScene) Draw(screen *ebiten.Image) {
	fontLarge := render.GetDefaultFont(render.FontLarge)
	fontMedium := render.GetDefaultFont(render.FontMedium)

	render.DrawText(screen, "Game Over", 5, 5, fontLarge)
	render.DrawText(screen, fmt.Sprintf("Score: %d", s.score), 5, 6, fontMedium)
	render.DrawText(screen, fmt.Sprintf("Level: %d", s.level), 5, 7, fontMedium)
	render.DrawText(screen, fmt.Sprintf("Lines: %d", s.lines), 5, 8, fontMedium)

	s.menu.Draw(screen, 5, 10)

	if s.isInitialsModeActive {
		render.DrawText(screen, fmt.Sprintf("Enter Initials: %s", s.initials), 5, 15, fontMedium)
		render.DrawText(screen, "Press ENTER to save", 5, 16, fontMedium)
		render.DrawText(screen, "Press ESC to cancel", 5, 17, fontMedium)
	}
}

func (s *GameOverScene) OnEnter() {}
func (s *GameOverScene) OnExit()  {}
