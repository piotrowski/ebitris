package scene

import "github.com/hajimehoshi/ebiten/v2"

type Manager struct {
	current Scene
	next    Scene
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) SwitchTo(scene Scene) {
	m.next = scene
}

func (m *Manager) Update() error {
	if m.next != nil {
		if m.current != nil {
			m.current.OnExit()
		}
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
