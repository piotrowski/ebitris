package tetris

import "math/rand/v2"

type Status int

const (
	StatusPlaying Status = iota
	StatusPaused
	StatusGameOver
)

type GameState struct {
	board        *Board
	currentPiece *Piece
	nextPiece    *Piece

	score        int
	linesCleared int

	status Status

	frameCount   int
	gravityDelay int // Frames between auto-drops
}

func (gs *GameState) GetLevel() int {
	return (gs.linesCleared / 10) + 1
}

func (gs *GameState) GetLinesCleared() int {
	return gs.linesCleared
}

func (gs *GameState) GetBoard() *Board {
	return gs.board
}

func (gs *GameState) GetCurrentPiece() *Piece {
	return gs.currentPiece
}

func (gs *GameState) GetNextPiece() *Piece {
	return gs.nextPiece
}

func (gs *GameState) GetScore() int {
	return gs.score
}

func (gs *GameState) Pause() {
	gs.status = StatusPaused
}

func (gs *GameState) Resume() {
	gs.status = StatusPlaying
}

func (gs *GameState) IsGameOver() bool {
	return gs.status == StatusGameOver
}

func NewGameState(width, height int) *GameState {
	return &GameState{
		board:        NewBoard(width, height),
		currentPiece: spawnRandomPiece(width/2-2, -2),
		nextPiece:    spawnRandomPiece(width/2-2, 0),
		gravityDelay: 48, // ~0.8 seconds at 60 FPS
		status:       StatusPlaying,
	}
}

var lastPieceID int

func pickRandomPiece() ShapeType {
	for i := 0; i < 2; i++ {
		id := rand.IntN(len(shapes))
		if id != lastPieceID {
			lastPieceID = id
			return ShapeType(id)
		}
	}
	return ShapeType(rand.IntN(len(shapes)))
}

func spawnRandomPiece(spawnX, spawnY int) *Piece {
	shape := pickRandomPiece()
	return NewPiece(shape, spawnX, spawnY, 0)
}

func (gs *GameState) Update() {
	if gs.status != StatusPlaying {
		return
	}

	gs.frameCount++
	if gs.frameCount >= gs.gravityDelay {
		gs.frameCount = 0
		gs.applyGravity()
	}
}

func (gs *GameState) MoveLeft() bool {
	if gs.board.IsColliding(gs.currentPiece, -1, 0) {
		return false
	}
	gs.currentPiece.MoveLeft()
	return true
}

func (gs *GameState) MoveRight() bool {
	if gs.board.IsColliding(gs.currentPiece, 1, 0) {
		return false
	}
	gs.currentPiece.MoveRight()
	return true
}

func (gs *GameState) MoveDown() bool {
	if gs.board.IsColliding(gs.currentPiece, 0, 1) {
		return false
	}
	gs.currentPiece.MoveDown()
	return true
}

func (gs *GameState) Rotate() bool {
	oldRotation := gs.currentPiece.Rotation
	gs.currentPiece.Rotate()
	if gs.board.IsColliding(gs.currentPiece, 0, 0) {
		gs.currentPiece.Rotation = oldRotation
		return false
	}
	return true
}

func (gs *GameState) HardDrop() {
	for !gs.board.IsColliding(gs.currentPiece, 0, 1) {
		gs.currentPiece.MoveDown()
	}
	gs.lockCurrentPiece()
}

func (gs *GameState) applyGravity() {
	// Try to move piece down
	if gs.board.IsColliding(gs.currentPiece, 0, 1) {
		gs.lockCurrentPiece()
	} else {
		gs.currentPiece.MoveDown()
	}
}

func (gs *GameState) lockCurrentPiece() {
	gs.board.LockPiece(gs.currentPiece)

	linesCleared := gs.board.ClearFullLines()
	if linesCleared > 0 {
		gs.addScore(linesCleared)
	}

	gs.currentPiece = gs.nextPiece
	gs.currentPiece.Y = -2
	gs.nextPiece = spawnRandomPiece(gs.board.Width/2-2, 0)

	if gs.board.IsGameOver() {
		gs.status = StatusGameOver
		return
	}
}

func (gs *GameState) addScore(linesCleared int) {
	points := map[int]int{
		1: 100, // Single
		2: 300, // Double
		3: 500, // Triple
		4: 800, // Tetris
	}

	currentLevel := gs.GetLevel()
	gs.score += points[linesCleared] * (currentLevel)
	gs.linesCleared += linesCleared

	newLevel := gs.GetLevel()
	if newLevel > currentLevel {
		gs.gravityDelay = max(5, 48-newLevel*3)
	}
}
