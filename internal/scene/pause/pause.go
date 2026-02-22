package pause

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/pkg/event"
	"github.com/piotrowski/ebitris/internal/pkg/input"
	"github.com/piotrowski/ebitris/internal/render"
)

type PauseScene struct {
	emitter event.Emitter
	input   *input.InputManager
	menu    *render.Menu
}

func NewPauseScene(emitter event.Emitter) *PauseScene {
	return &PauseScene{
		emitter: emitter,
		input:   input.NewInputManager(),
		menu:    render.NewMenu([]string{"Resume", "Restart", "Main Menu"}),
	}
}

func (s *PauseScene) Update() error {
	if s.input.IsKeyJustPressed(ebiten.KeyEscape) {
		s.emitter.Emit(event.Event{Type: event.EventTypeGoBack})
	}

	if s.menu.HandleInput(s.input) {
		switch s.menu.Selected() {
		case 0:
			s.emitter.Emit(event.Event{Type: event.EventTypeGoBack})
		case 1:
			s.emitter.Emit(event.Event{Type: event.EventTypeStartGame})
		case 2:
			s.emitter.Emit(event.Event{Type: event.EventTypeMainMenu})
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
