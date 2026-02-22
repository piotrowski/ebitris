package scoreboard

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/pkg/event"
	"github.com/piotrowski/ebitris/internal/pkg/input"
	"github.com/piotrowski/ebitris/internal/pkg/persistence"
	"github.com/piotrowski/ebitris/internal/render"
)

var pageSize = 10

type ScoreboardScene struct {
	emitter event.Emitter
	input   *input.InputManager
	menu    *render.Menu

	// scoreManager ScoreManager

	currentPage  int
	hasMorePages bool
	scores       []persistence.ScoreEntry
}

func NewScoreboardScene(emitter event.Emitter) *ScoreboardScene {
	s := &ScoreboardScene{
		emitter: emitter,
		input:   input.NewInputManager(),
		menu:    render.NewMenu([]string{"Next Page", "Previous Page", "Back"}),
	}

	// s.scores, s.hasMorePages = s.scoreManager.GetPage(s.currentPage, pageSize)
	return s
}

func (s *ScoreboardScene) Update() error {
	if s.menu.HandleInput(s.input) {
		switch s.menu.Selected() {
		case 0:
			if s.hasMorePages {
				s.currentPage++
				// s.scores, s.hasMorePages = s.scoreManager.GetPage(s.currentPage, pageSize)
			}
		case 1:
			if s.currentPage > 0 {
				s.currentPage--
			}
			// s.scores, s.hasMorePages = s.scoreManager.GetPage(s.currentPage, pageSize)
		case 2:
			s.emitter.Emit(event.Event{Type: event.EventTypeMainMenu})
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
