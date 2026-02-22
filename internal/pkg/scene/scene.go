package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	OnEnter()
	OnExit()
}

type Switcher interface {
	SwitchTo(scene Scene)
	SwitchBack()
}

type Quitter interface {
	Quit()
}

type Updater interface {
	Update() error
}

type Drawer interface {
	Draw(screen *ebiten.Image)
}

type Manager interface {
	Switcher
	Quitter
	Updater
	Drawer
}

type SceneManager struct {
	prev    Scene
	current Scene
	next    Scene

	quit bool
}

func NewSceneManager() *SceneManager {
	return &SceneManager{}
}

func (m *SceneManager) SwitchTo(scene Scene) {
	m.next = scene
}

func (m *SceneManager) SwitchBack() {
	if m.prev != nil {
		m.next = m.prev
	}
}

func (m *SceneManager) Update() error {
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

func (m *SceneManager) Draw(screen *ebiten.Image) {
	if m.current != nil {
		m.current.Draw(screen)
	}
}

func (m *SceneManager) Quit() {
	m.quit = true
}
