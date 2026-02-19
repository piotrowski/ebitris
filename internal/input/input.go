package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputManager struct {
	keyHoldTime map[ebiten.Key]int
}

const (
	InitialDelay = 15 // Frames before repeat starts
	RepeatDelay  = 3  // Frames between repeats
)

func NewInputManager() *InputManager {
	return &InputManager{
		keyHoldTime: make(map[ebiten.Key]int),
	}
}

func (im *InputManager) IsKeyJustPressed(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key)
}

func (im *InputManager) IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

func (im *InputManager) ShouldMove(key ebiten.Key) bool {
	if im.IsKeyJustPressed(key) {
		im.keyHoldTime[key] = 0
		return true
	}

	if im.IsKeyPressed(key) {
		im.keyHoldTime[key]++
		if im.keyHoldTime[key] >= InitialDelay {
			if (im.keyHoldTime[key]-InitialDelay)%RepeatDelay == 0 {
				return true
			}
		}
	} else {
		im.keyHoldTime[key] = 0
	}

	return false
}
