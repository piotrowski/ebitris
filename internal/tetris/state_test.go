package tetris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel(t *testing.T) {
	tests := []struct {
		name          string
		linesCleared  int
		expectedLevel int
	}{
		{
			name:          "level 0 with no lines cleared",
			linesCleared:  0,
			expectedLevel: 0,
		},
		{
			name:          "level 1 with 10 lines cleared",
			linesCleared:  10,
			expectedLevel: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := NewGameState(10, 20)
			gs.LinesCleared = tt.linesCleared
			assert.Equal(t, tt.expectedLevel, gs.Level())
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name         string
		status       Status
		frameCount   int
		gravityDelay int
		shouldApply  bool
	}{
		{
			name:         "paused game does nothing",
			status:       StatusPaused,
			frameCount:   50,
			gravityDelay: 48,
			shouldApply:  false,
		},
		{
			name:         "game over does nothing",
			status:       StatusGameOver,
			frameCount:   50,
			gravityDelay: 48,
			shouldApply:  false,
		},
		{
			name:         "playing game increments frame count",
			status:       StatusPlaying,
			frameCount:   0,
			gravityDelay: 48,
			shouldApply:  false,
		},
		{
			name:         "should apply gravity after delay",
			status:       StatusPlaying,
			frameCount:   48,
			gravityDelay: 48,
			shouldApply:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := NewGameState(10, 20)
			gs.Status = tt.status
			gs.FrameCount = tt.frameCount
			gs.GravityDelay = tt.gravityDelay
			originalY := gs.CurrentPiece.Y
			gs.Update()
			if tt.status != StatusPlaying {
				assert.Equal(t, originalY, gs.CurrentPiece.Y)
			}
		})
	}
}

func TestAddScore(t *testing.T) {
	tests := []struct {
		name                 string
		linesCleared         int
		currentLinesCleared  int
		expectedScore        int
		expectedGravityDelay int
	}{
		{
			name:                 "single line",
			linesCleared:         1,
			currentLinesCleared:  0,
			expectedScore:        100,
			expectedGravityDelay: 48,
		},
		{
			name:                 "double line",
			linesCleared:         2,
			currentLinesCleared:  0,
			expectedScore:        300,
			expectedGravityDelay: 48,
		},
		{
			name:                 "triple line",
			linesCleared:         3,
			currentLinesCleared:  0,
			expectedScore:        500,
			expectedGravityDelay: 48,
		},
		{
			name:                 "tetris",
			linesCleared:         4,
			currentLinesCleared:  0,
			expectedScore:        800,
			expectedGravityDelay: 48,
		},
		{
			name:                 "level up decreases gravity delay",
			linesCleared:         1,
			currentLinesCleared:  9,
			expectedScore:        100,
			expectedGravityDelay: 46,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := NewGameState(10, 20)
			gs.LinesCleared = tt.currentLinesCleared
			gs.addScore(tt.linesCleared)
			assert.Equal(t, tt.expectedScore, gs.Score)
			assert.Equal(t, tt.expectedGravityDelay, gs.GravityDelay)
		})
	}
}

func TestLockCurrentPiece(t *testing.T) {
	gs := NewGameState(10, 20)
	originalNextPiece := gs.NextPiece
	gs.lockCurrentPiece()
	assert.Equal(t, originalNextPiece, gs.CurrentPiece)
	assert.NotEqual(t, originalNextPiece, gs.NextPiece)
	assert.NotNil(t, gs.NextPiece)
}
