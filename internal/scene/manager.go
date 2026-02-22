package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/music"
	"github.com/piotrowski/ebitris/internal/persistence"
)

type ScoreManager interface {
	SaveScore(initials string, score, level, lines int)
	GetPage(page int, size int) ([]persistence.ScoreEntry, bool)
}

type AudioManager interface {
	PlaySong(name music.SongName)
	StartAutoPlay()
	StopAutoPlay()
	SetPlaylist(names ...music.SongName)

	PlayEffect(name music.EffectName)
}

type Manager struct {
	prev    Scene
	current Scene
	next    Scene

	quit         bool
	scoreManager ScoreManager
	audioManager AudioManager
}

func NewManager() *Manager {
	return &Manager{
		scoreManager: persistence.NewScoreManager(),
		audioManager: music.NewManager(),
	}
}

func (m *Manager) SwitchTo(scene Scene) {
	m.next = scene
}

func (m *Manager) SwitchBack() {
	if m.prev != nil {
		m.next = m.prev
	}
}

func (m *Manager) Update() error {
	if m.quit {
		return ebiten.Termination
	}

	if m.next != nil {
		if m.current != nil {
			m.current.OnExit()
		}
		m.prev = m.current
		m.current = m.next
		m.current.OnEnter()
		m.next = nil
	}

	if m.current != nil {
		return m.current.Update()
	}
	return nil
}

func (m *Manager) Draw(screen *ebiten.Image) {
	if m.current != nil {
		m.current.Draw(screen)
	}
}

func (m *Manager) Quit() {
	m.quit = true
}
