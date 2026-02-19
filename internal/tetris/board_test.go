package tetris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsColliding(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		board    *Board
		piece    *Piece
		expected bool
	}{
		{
			name:     "no collision with empty board",
			board:    NewBoard(10, 20),
			piece:    &Piece{X: 5, Y: 5, Shape: ShapeI},
			expected: false,
		},
		{
			name:     "collision with left boundary",
			board:    NewBoard(10, 20),
			piece:    &Piece{X: -1, Y: 5, Shape: ShapeI},
			expected: true,
		},
		{
			name:     "collision with right boundary",
			board:    NewBoard(10, 20),
			piece:    &Piece{X: 10, Y: 5, Shape: ShapeI},
			expected: true,
		},
		{
			name:     "collision with top boundary",
			board:    NewBoard(10, 20),
			piece:    &Piece{X: 5, Y: -1, Shape: ShapeI},
			expected: false,
		},
		{
			name:     "collision with occupied cell",
			board:    func() *Board { b := NewBoard(10, 20); b.grid[6][5] = 1; return b }(),
			piece:    &Piece{X: 5, Y: 5, Shape: ShapeI},
			expected: true,
		},
		{
			name:     "no collision with multiple cells",
			board:    NewBoard(10, 20),
			piece:    &Piece{X: 3, Y: 3, Shape: ShapeI},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.board.IsColliding(tt.piece, 0, 0)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLockPiece(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		board         *Board
		piece         *Piece
		expectedGrid  [][]int  // key: "x,y", value: color
		expectedEmpty []string // cells that should remain 0
	}{
		{
			name:  "lock piece in empty area",
			board: NewBoard(5, 5),
			piece: &Piece{X: 0, Y: 0, Shape: ShapeI, Color: 1},
			expectedGrid: [][]int{
				{0, 0, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.board.LockPiece(tt.piece)
			for y, row := range tt.expectedGrid {
				for x, expectedColor := range row {
					assert.Equal(t, expectedColor, tt.board.grid[y][x])
				}
			}
		})
	}
}

func TestClearFullLines(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		board           *Board
		expectedCleared int
		expectedGrid    [][]int
	}{
		{
			name: "no full lines",
			board: func() *Board {
				b := NewBoard(5, 5)
				b.grid[4][0] = 1
				b.grid[4][1] = 1
				return b
			}(),
			expectedCleared: 0,
			expectedGrid: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
			},
		},
		{
			name: "single full line at bottom",
			board: func() *Board {
				b := NewBoard(5, 5)
				b.grid[4][0] = 1
				b.grid[4][1] = 1
				b.grid[4][2] = 1
				b.grid[4][3] = 1
				b.grid[4][4] = 1
				return b
			}(),
			expectedCleared: 1,
			expectedGrid: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
		},
		{
			name: "full line with data above",
			board: func() *Board {
				b := NewBoard(3, 4)
				b.grid[1][0] = 1
				b.grid[2][0] = 1
				b.grid[2][1] = 1
				b.grid[2][2] = 1
				b.grid[3][0] = 2
				b.grid[3][1] = 2
				b.grid[3][2] = 2
				return b
			}(),
			expectedCleared: 2,
			expectedGrid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
				{1, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cleared := tt.board.ClearFullLines()
			assert.Equal(t, tt.expectedCleared, cleared)
			for y, row := range tt.expectedGrid {
				for x, expectedValue := range row {
					assert.Equal(t, expectedValue, tt.board.grid[y][x])
				}
			}
		})
	}
}
