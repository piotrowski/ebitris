package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/pkg/input"
)

type Menu struct {
	items []string
	focus int
}

func NewMenu(items []string) *Menu {
	return &Menu{items: items}
}

// HandleInput handles Up/Down/Enter navigation. Returns true if Enter was pressed.
func (m *Menu) HandleInput(im *input.InputManager) bool {
	n := len(m.items)
	if im.IsKeyJustPressed(ebiten.KeyDown) {
		m.focus = (m.focus + 1) % n
	}
	if im.IsKeyJustPressed(ebiten.KeyUp) {
		m.focus = (m.focus - 1 + n) % n
	}
	return im.IsKeyJustPressed(ebiten.KeyEnter)
}

// Selected returns the index of the currently focused item.
func (m *Menu) Selected() int {
	return m.focus
}

// Draw renders the selection indicator and item labels starting at cell (x, y).
func (m *Menu) Draw(screen *ebiten.Image, x, y int) {
	DrawRectangle(screen, x-1, y+m.focus, 5, 1, color.RGBA{16, 16, 16, 255})

	fontMedium := GetDefaultFont(FontMedium)
	for i, label := range m.items {
		DrawText(screen, label, x, y+i, fontMedium)
	}
}
