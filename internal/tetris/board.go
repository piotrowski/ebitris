package tetris

type Board struct {
	Width  int
	Height int
	grid   [][]int // 0 for empty, 1-7 for different colors
}

func NewBoard(width, height int) *Board {
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	return &Board{
		Width:  width,
		Height: height,
		grid:   grid,
	}
}

func (b *Board) Cell(x, y int) int {
	return b.grid[y][x]
}

func (b *Board) IsColliding(piece *Piece, offsetX, offsetY int) bool {
	for _, cell := range piece.GetCells() {
		x := piece.X + cell.X + offsetX
		y := piece.Y + cell.Y + offsetY

		if x < 0 || x >= b.Width {
			return true
		}

		if y >= b.Height {
			return true
		}

		if y >= 0 && b.grid[y][x] != 0 {
			return true
		}
	}
	return false
}

func (b *Board) LockPiece(piece *Piece) {
	for _, cell := range piece.GetCells() {
		x := piece.X + cell.X
		y := piece.Y + cell.Y
		if x >= 0 && x < b.Width && y >= 0 && y < b.Height {
			b.grid[y][x] = int(piece.Color)
		}
	}
}

func (b *Board) ClearFullLines() int {
	linesCleared := 0

	for y := b.Height - 1; y >= 0; y-- {
		if b.isLineFull(y) {
			b.removeLine(y)
			linesCleared++
			y++ // Check the same line again after shifting down
		}
	}

	return linesCleared
}

func (b *Board) isLineFull(y int) bool {
	for x := 0; x < b.Width; x++ {
		if b.grid[y][x] == 0 {
			return false
		}
	}
	return true
}

func (b *Board) removeLine(lineY int) {
	for y := lineY; y > 0; y-- {
		b.grid[y] = b.grid[y-1]
	}
	b.grid[0] = make([]int, b.Width)
}

// IsGameOver checks if any cells in the top row are occupied, which indicates game over.
func (b *Board) IsGameOver() bool {
	for x := 0; x < b.Width; x++ {
		if b.grid[0][x] != 0 {
			return true
		}
	}
	return false
}
