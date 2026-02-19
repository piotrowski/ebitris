package tetris

import "math/rand/v2"

type Status int

const (
	StatusPlaying Status = iota
	StatusPaused
	StatusGameOver
)

type GameState struct {
	Board        *Board
	CurrentPiece *Piece
	NextPiece    *Piece

	Score        int
	LinesCleared int

	status Status

	FrameCount   int
	GravityDelay int // Frames between auto-drops
}

func (gs *GameState) Level() int {
	return gs.LinesCleared / 10
}

func (gs *GameState) Pause() {
	gs.status = StatusPaused
}

func (gs *GameState) Resume() {
	gs.status = StatusPlaying
}

func NewGameState(width, height int) *GameState {
	return &GameState{
		Board:        NewBoard(width, height),
		CurrentPiece: spawnRandomPiece(width/2-2, 0),
		NextPiece:    spawnRandomPiece(width/2-2, 0),
		GravityDelay: 48, // ~0.8 seconds at 60 FPS
		status:       StatusPlaying,
	}
}

func spawnRandomPiece(spawnX, spawnY int) *Piece {
	shape := ShapeType(rand.IntN(7))
	return NewPiece(shape, spawnX, spawnY, 0) // Start near the top center
}

func (gs *GameState) Update() {
	if gs.status != StatusPlaying {
		return
	}

	// Apply gravity
	gs.FrameCount++
	if gs.FrameCount >= gs.GravityDelay {
		gs.FrameCount = 0
		gs.applyGravity()
	}
}

func (gs *GameState) applyGravity() {
	// Try to move piece down
	if gs.Board.IsColliding(gs.CurrentPiece, 0, 1) {
		gs.lockCurrentPiece()
	} else {
		gs.CurrentPiece.MoveDown()
	}
}

func (gs *GameState) lockCurrentPiece() {
	gs.Board.LockPiece(gs.CurrentPiece)

	// Clear lines
	linesCleared := gs.Board.ClearFullLines()
	if linesCleared > 0 {
		gs.addScore(linesCleared)
	}

	// Spawn next piece
	gs.CurrentPiece = gs.NextPiece
	gs.NextPiece = spawnRandomPiece(gs.Board.Width/2-2, 0)

	// Check game over
	if gs.Board.IsGameOver() {
		gs.status = StatusGameOver
	}
}

func (gs *GameState) addScore(linesCleared int) {
	points := map[int]int{
		1: 100, // Single
		2: 300, // Double
		3: 500, // Triple
		4: 800, // Tetris
	}

	currentLevel := gs.Level()
	gs.Score += points[linesCleared] * (currentLevel + 1)
	gs.LinesCleared += linesCleared

	newLevel := gs.Level()
	if newLevel > currentLevel {
		gs.GravityDelay = max(10, 48-newLevel*2)
	}
}
