package scene

import (
	"errors"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

type mockScene struct {
	enterCount  int
	exitCount   int
	updateCount int
	updateErr   error
}

func (m *mockScene) Update() error {
	m.updateCount++
	return m.updateErr
}

func (m *mockScene) Draw(_ *ebiten.Image) {}

func (m *mockScene) OnEnter() { m.enterCount++ }
func (m *mockScene) OnExit()  { m.exitCount++ }

func TestSceneManager(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		run  func(t *testing.T, m *SceneManager)
	}{
		{
			name: "SwitchTo triggers lifecycle",
			run: func(t *testing.T, m *SceneManager) {
				t.Helper()
				a := &mockScene{}
				b := &mockScene{}
				m.current = a

				m.SwitchTo(b)
				err := m.Update()

				assert.NoError(t, err)
				assert.Equal(t, 1, a.exitCount)
				assert.Equal(t, 1, b.enterCount)
				assert.Equal(t, b, m.current)
			},
		},
		{
			name: "SwitchTo with no current scene",
			run: func(t *testing.T, m *SceneManager) {
				t.Helper()
				a := &mockScene{}

				assert.NotPanics(t, func() {
					m.SwitchTo(a)
					_ = m.Update()
				})
				assert.Equal(t, 1, a.enterCount)
			},
		},
		{
			name: "SwitchBack restores previous scene",
			run: func(t *testing.T, m *SceneManager) {
				t.Helper()
				a := &mockScene{}
				b := &mockScene{}

				m.SwitchTo(a)
				_ = m.Update()

				m.SwitchTo(b)
				_ = m.Update()

				m.SwitchBack()
				_ = m.Update()

				assert.Equal(t, 2, a.enterCount)
				assert.Equal(t, a, m.current)
			},
		},
		{
			name: "SwitchBack with no previous scene",
			run: func(t *testing.T, m *SceneManager) {
				t.Helper()
				assert.NotPanics(t, func() {
					m.SwitchBack()
					_ = m.Update()
				})
				assert.Nil(t, m.current)
			},
		},
		{
			name: "Quit returns ebiten.Termination",
			run: func(t *testing.T, m *SceneManager) {
				t.Helper()
				m.Quit()
				err := m.Update()
				assert.Equal(t, ebiten.Termination, err)
			},
		},
		{
			name: "Update with nil current returns nil",
			run: func(t *testing.T, m *SceneManager) {
				t.Helper()
				err := m.Update()
				assert.NoError(t, err)
			},
		},
		{
			name: "transition applies once across multiple Updates",
			run: func(t *testing.T, m *SceneManager) {
				t.Helper()
				a := &mockScene{}

				m.SwitchTo(a)
				_ = m.Update()
				_ = m.Update()

				assert.Equal(t, 1, a.enterCount)
			},
		},
		{
			name: "Update propagates scene error",
			run: func(t *testing.T, m *SceneManager) {
				t.Helper()
				boom := errors.New("boom")
				a := &mockScene{updateErr: boom}
				m.current = a

				err := m.Update()
				assert.Equal(t, boom, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := NewSceneManager()
			tt.run(t, m)
		})
	}
}
