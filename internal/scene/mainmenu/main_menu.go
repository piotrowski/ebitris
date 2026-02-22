package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/pkg/event"
	"github.com/piotrowski/ebitris/internal/pkg/input"
	"github.com/piotrowski/ebitris/internal/render"
)

type MenuScene struct {
	emitter event.Emitter

	input *input.InputManager
	menu  *render.Menu
}

func NewMenuScene(emitter event.Emitter) *MenuScene {
	return &MenuScene{
		emitter: emitter,
		input:   input.NewInputManager(),
		menu:    render.NewMenu([]string{"Start Game", "Scoreboard", "Exit"}),
	}
}

func (s *MenuScene) Update() error {
	if s.menu.HandleInput(s.input) {
		switch s.menu.Selected() {
		case 0:
			s.emitter.Emit(event.Event{Type: event.EventTypeStartGame})
		case 1:
			s.emitter.Emit(event.Event{Type: event.EventTypeScoreboard})
		case 2:
			s.emitter.Emit(event.Event{Type: event.EventTypeQuit})
		}
	}

	return nil
}

func (s *MenuScene) Draw(screen *ebiten.Image) {
	fontLarge := render.GetDefaultFont(render.FontLarge)
	render.DrawText(screen, "Ebitris", 5, 5, fontLarge)
	s.menu.Draw(screen, 5, 10)
}

func (s *MenuScene) OnEnter() {
	// s.manager.audioManager.PlaySong(music.ReturnOfThe8BitEra)
}
func (s *MenuScene) OnExit() {}
