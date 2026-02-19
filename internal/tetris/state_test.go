package tetris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			gs := NewGameState(10, 20)
			gs.linesCleared = tt.linesCleared
			assert.Equal(t, tt.expectedLevel, gs.GetLevel())
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			gs := NewGameState(10, 20)
			gs.status = tt.status
			gs.frameCount = tt.frameCount
			gs.gravityDelay = tt.gravityDelay
			originalY := gs.currentPiece.Y
			gs.Update()
			if tt.status != StatusPlaying {
				assert.Equal(t, originalY, gs.currentPiece.Y)
			}
		})
	}
}

func TestAddScore(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			gs := NewGameState(10, 20)
			gs.linesCleared = tt.currentLinesCleared
			gs.addScore(tt.linesCleared)
			assert.Equal(t, tt.expectedScore, gs.score)
			assert.Equal(t, tt.expectedGravityDelay, gs.gravityDelay)
		})
	}
}

func TestMoveLeft(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		startX   int
		expected bool
	}{
		{
			name:     "moves left when space available",
			startX:   5,
			expected: true,
		},
		{
			name:     "blocked at left wall",
			startX:   0,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gs := NewGameState(10, 20)
			gs.currentPiece = NewPiece(ShapeO, tt.startX, 0, 0)
			result := gs.MoveLeft()
			assert.Equal(t, tt.expected, result)
			if tt.expected {
				assert.Equal(t, tt.startX-1, gs.currentPiece.X)
			} else {
				assert.Equal(t, tt.startX, gs.currentPiece.X)
			}
		})
	}
}

func TestMoveRight(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		startX   int
		expected bool
	}{
		{
			name:     "moves right when space available",
			startX:   0,
			expected: true,
		},
		{
			name:     "blocked at right wall",
			startX:   8,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gs := NewGameState(10, 20)
			gs.currentPiece = NewPiece(ShapeO, tt.startX, 0, 0)
			result := gs.MoveRight()
			assert.Equal(t, tt.expected, result)
			if tt.expected {
				assert.Equal(t, tt.startX+1, gs.currentPiece.X)
			} else {
				assert.Equal(t, tt.startX, gs.currentPiece.X)
			}
		})
	}
}

func TestMoveDown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		startY   int
		expected bool
	}{
		{
			name:     "moves down when space available",
			startY:   0,
			expected: true,
		},
		{
			name:     "blocked at bottom",
			startY:   18,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gs := NewGameState(10, 20)
			gs.currentPiece = NewPiece(ShapeO, 4, tt.startY, 0)
			result := gs.MoveDown()
			assert.Equal(t, tt.expected, result)
			if tt.expected {
				assert.Equal(t, tt.startY+1, gs.currentPiece.Y)
			} else {
				assert.Equal(t, tt.startY, gs.currentPiece.Y)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		shape           ShapeType
		startX          int
		startY          int
		startRot        int
		expectedSuccess bool
		expectedRot     int
	}{
		{
			name:            "rotates when space available",
			shape:           ShapeT,
			startX:          4,
			startY:          1,
			expectedSuccess: true,
			expectedRot:     1,
		},
		{
			name:            "reverts rotation on top-wall collision",
			shape:           ShapeL,
			startX:          -1,
			startY:          -1,
			startRot:        1,
			expectedSuccess: false,
			expectedRot:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gs := NewGameState(10, 20)
			gs.currentPiece = NewPiece(tt.shape, tt.startX, tt.startY, tt.startRot)
			result := gs.Rotate()
			assert.Equal(t, tt.expectedSuccess, result)
			assert.Equal(t, tt.expectedRot, gs.currentPiece.Rotation)
		})
	}
}

func TestHardDrop(t *testing.T) {
	t.Parallel()

	gs := NewGameState(10, 20)
	gs.currentPiece = NewPiece(ShapeO, 4, 0, 0)
	gs.HardDrop()

	// ShapeO cells at relative Y=1,2; board height=20.
	// Lowest valid pieceY: 17 (abs rows 18,19 are in bounds; abs row 20 would be out).
	assert.Equal(t, 18, gs.currentPiece.Y)
	// Moving down further should be blocked
	assert.False(t, gs.MoveDown())
}

func TestLockCurrentPiece(t *testing.T) {
	t.Parallel()

	gs := NewGameState(10, 20)
	originalNextPiece := gs.nextPiece
	gs.lockCurrentPiece()
	assert.Equal(t, originalNextPiece, gs.currentPiece)
	assert.NotEqual(t, originalNextPiece, gs.nextPiece)
	assert.NotNil(t, gs.nextPiece)
}
