package scene

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/input"
	"github.com/piotrowski/ebitris/internal/persistence"
	"github.com/piotrowski/ebitris/internal/render"
	"github.com/piotrowski/ebitris/internal/ui"
)

var pageSize = 10

type ScoreboardScene struct {
	manager *Manager
	input   *input.InputManager
	menu    *ui.Menu

	scoreManager ScoreManager

	currentPage  int
	hasMorePages bool
	scores       []persistence.ScoreEntry
}

func NewScoreboardScene(manager *Manager) *ScoreboardScene {
	s := &ScoreboardScene{
		manager: manager,
		input:   input.NewInputManager(),
		menu:    ui.NewMenu([]string{"Next Page", "Previous Page", "Back"}),

		scoreManager: manager.scoreManager,
	}

	s.scores, s.hasMorePages = s.scoreManager.GetPage(s.currentPage, pageSize)
	return s
}

func (s *ScoreboardScene) Update() error {
	if s.menu.HandleInput(s.input) {
		switch s.menu.Selected() {
		case 0:
			if s.hasMorePages {
				s.currentPage++
				s.scores, s.hasMorePages = s.scoreManager.GetPage(s.currentPage, pageSize)
			}
		case 1:
			if s.currentPage > 0 {
				s.currentPage--
			}
			s.scores, s.hasMorePages = s.scoreManager.GetPage(s.currentPage, pageSize)
		case 2:
			s.manager.SwitchTo(NewMainMenuScene(s.manager))
		}
	}

	return nil
}

func (s *ScoreboardScene) Draw(screen *ebiten.Image) {
	fontLarge := render.GetDefaultFont(render.FontLarge)
	fontMedium := render.GetDefaultFont(render.FontMedium)
	render.DrawText(screen, "Ebitris", 5, 5, fontLarge)
	s.menu.Draw(screen, 5, 10)

	for i, score := range s.scores {
		render.DrawText(
			screen,
			fmt.Sprintf("%d. %s - Score: %d, Level: %d, Lines: %d", s.currentPage*pageSize+i+1, score.Initials, score.Score, score.Level, score.Lines),
			5,
			15+i,
			fontMedium,
		)
	}
}

func (s *ScoreboardScene) OnEnter() {}
func (s *ScoreboardScene) OnExit()  {}
