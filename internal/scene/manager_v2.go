package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/pkg/event"
	"github.com/piotrowski/ebitris/internal/pkg/scene"
	"github.com/piotrowski/ebitris/internal/pkg/score"
	"github.com/piotrowski/ebitris/internal/scene/gameover"
	"github.com/piotrowski/ebitris/internal/scene/gameplay"
	menu "github.com/piotrowski/ebitris/internal/scene/mainmenu"
	"github.com/piotrowski/ebitris/internal/scene/pause"
	"github.com/piotrowski/ebitris/internal/scene/scoreboard"
)

type eventManager interface {
	event.Emitter
	event.Subscriber
	event.Dispatcher
}

type ManagerV2 struct {
	events       eventManager
	sceneManager scene.Manager
	scoreManager score.Saver
}

func NewManagerV2() *ManagerV2 {
	m := &ManagerV2{
		events:       event.NewEventManager(),
		sceneManager: scene.NewSceneManager(),
		scoreManager: score.NewScoreManager(),
	}

	m.subscribeNavigation()
	m.sceneManager.SwitchTo(menu.NewMenuScene(m.events))

	return m
}

func (m *ManagerV2) subscribeNavigation() {
	m.events.Subscribe(event.EventTypeGoBack, func(e event.Event) {
		m.sceneManager.SwitchBack()
	})

	m.events.Subscribe(event.EventTypeStartGame, func(e event.Event) {
		m.sceneManager.SwitchTo(gameplay.NewStandardGameplayScene(m.events))
	})

	m.events.Subscribe(event.EventTypeMainMenu, func(e event.Event) {
		m.sceneManager.SwitchTo(menu.NewMenuScene(m.events))
	})

	m.events.Subscribe(event.EventTypeScoreboard, func(e event.Event) {
		m.sceneManager.SwitchTo(scoreboard.NewScoreboardScene(m.events))
	})

	m.events.Subscribe(event.EventTypePause, func(e event.Event) {
		m.sceneManager.SwitchTo(pause.NewPauseScene(m.events))
	})

	m.events.Subscribe(event.EventTypeGameOver, func(e event.Event) {
		m.sceneManager.SwitchTo(gameover.NewGameOverScene(m.events, 0, 0, 0))
	})

	m.events.Subscribe(event.EventTypeQuit, func(e event.Event) {
		m.sceneManager.Quit()
	})
}

func (m *ManagerV2) Update() error {
	err := m.sceneManager.Update()
	m.events.Dispatch()
	return err
}

func (m *ManagerV2) Draw(screen *ebiten.Image) {
	m.sceneManager.Draw(screen)
}
