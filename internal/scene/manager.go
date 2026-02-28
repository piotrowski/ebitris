package scene

import (
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/pkg/audio"
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

type scoreManager interface {
	score.Getter
	score.Saver
}

type audioManager interface {
	audio.EffectPlayer
	audio.MusicPlayer
	audio.AudioUpdater
}

type Manager struct {
	events       eventManager
	sceneManager scene.Manager
	scoreManager scoreManager
	audioManager audioManager
}

func NewManager() *Manager {
	m := &Manager{
		events:       event.NewEventManager(),
		sceneManager: scene.NewSceneManager(),
		scoreManager: score.NewScoreManager(),
		audioManager: audio.NewAudioManager(),
	}

	m.subscribeNavigation()
	m.subscribeMusic()
	m.subscribeEffects()

	m.events.Emit(event.Event{Type: event.EventTypeMainMenu})

	return m
}

func (m *Manager) subscribeNavigation() {
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
		m.sceneManager.SwitchTo(scoreboard.NewScoreboardScene(m.events, m.scoreManager))
	})

	m.events.Subscribe(event.EventTypePause, func(e event.Event) {
		m.sceneManager.SwitchTo(pause.NewPauseScene(m.events))
	})

	m.events.Subscribe(event.EventTypeGameOver, func(e event.Event) {
		endScore, isOk := e.Payload.(event.GameOverPayload)
		if !isOk {
			slog.Warn("unexpected GameOverPayload", "subsystem", "scene")
		}
		m.sceneManager.SwitchTo(gameover.NewGameOverScene(m.events, m.scoreManager, endScore.Score, endScore.Level, endScore.Lines))
	})

	m.events.Subscribe(event.EventTypeQuit, func(e event.Event) {
		m.sceneManager.Quit()
	})
}

func (m *Manager) subscribeMusic() {
	m.events.Subscribe(event.EventTypeStartGame, func(e event.Event) {
		m.audioManager.StartPlaylist(audio.ArcadeBeat, audio.ReturnOfThe8BitEra)
	})

	m.events.Subscribe(event.EventTypeMainMenu, func(e event.Event) {
		m.audioManager.StartPlaylist(audio.ArcadeBeat, audio.ReturnOfThe8BitEra)
	})

	m.events.Subscribe(event.EventTypeGameOver, func(e event.Event) {
		m.audioManager.StartPlaylist(audio.ArcadeBeat, audio.ReturnOfThe8BitEra)
	})
}

func (m *Manager) subscribeEffects() {
	m.events.Subscribe(event.EventTypeBlockPlaced, func(e event.Event) {
		m.audioManager.PlayEffect(audio.ExplosionEffect)
	})

	m.events.Subscribe(event.EventTypeBlockMovedByPlayer, func(e event.Event) {
		m.audioManager.PlayEffect(audio.BipEffect)
	})
}

func (m *Manager) Update() error {
	err := m.sceneManager.Update()
	if err != nil {
		return err
	}

	err = m.audioManager.Update()
	if err != nil {
		return err
	}

	m.events.Dispatch()
	return nil
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.sceneManager.Draw(screen)
}
