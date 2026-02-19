package scene

import "github.com/hajimehoshi/ebiten/v2"

type Manager struct {
	prev    Scene
	current Scene
	next    Scene
}

func NewManager() *Manager {
	return &Manager{}
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
