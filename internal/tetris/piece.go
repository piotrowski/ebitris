package tetris

import "image/color"

type PieceColor int

const (
	PieceRed PieceColor = iota + 1
	PieceGreen
	PieceBlue
	PieceYellow
	PieceCyan
	PieceMagenta
	PieceOrange
	PieceShadow
)

var shapeColors = map[ShapeType]PieceColor{
	ShapeI: PieceCyan,
	ShapeO: PieceYellow,
	ShapeT: PieceMagenta,
	ShapeS: PieceGreen,
	ShapeZ: PieceRed,
	ShapeJ: PieceBlue,
	ShapeL: PieceOrange,
}

var PieceColors = map[PieceColor]color.Color{
	PieceRed:     color.RGBA{R: 255, G: 0, B: 0, A: 255},
	PieceGreen:   color.RGBA{R: 0, G: 255, B: 0, A: 255},
	PieceBlue:    color.RGBA{R: 0, G: 0, B: 255, A: 255},
	PieceYellow:  color.RGBA{R: 255, G: 255, B: 0, A: 255},
	PieceCyan:    color.RGBA{R: 0, G: 255, B: 255, A: 255},
	PieceMagenta: color.RGBA{R: 255, G: 0, B: 255, A: 255},
	PieceOrange:  color.RGBA{R: 255, G: 165, B: 0, A: 255},
	PieceShadow:  color.RGBA{R: 40, G: 40, B: 50, A: 255},
}

func GetPieceColor[T ~int](c T) color.Color {
	if col, ok := PieceColors[PieceColor(c)]; ok {
		return col
	}
	return color.RGBA{R: 255, G: 255, B: 255, A: 255} // default to white for invalid colors
}

type Piece struct {
	Shape    ShapeType
	Color    PieceColor
	X        int
	Y        int
	Rotation int
}

type Cell struct {
	X, Y int
}

func NewPiece(shape ShapeType, posX, posY, rotation int) *Piece {
	return &Piece{
		Shape:    shape,
		Color:    shapeColors[shape],
		X:        posX,
		Y:        posY,
		Rotation: rotation,
	}
}

func (p *Piece) Clone() *Piece {
	return &Piece{
		Shape:    p.Shape,
		Color:    p.Color,
		X:        p.X,
		Y:        p.Y,
		Rotation: p.Rotation,
	}
}

func (p *Piece) GetCells() []Cell {
	return shapes[p.Shape][p.Rotation]
}

func (p *Piece) Rotate() {
	p.Rotation = (p.Rotation + 1) % len(shapes[p.Shape])
}

func (p *Piece) MoveLeft() {
	p.move(-1, 0)
}

func (p *Piece) MoveRight() {
	p.move(1, 0)
}

func (p *Piece) MoveDown() {
	p.move(0, 1)
}

func (p *Piece) MoveUp() {
	p.move(0, -1)
}

func (p *Piece) move(dx, dy int) {
	p.X += dx
	p.Y += dy
}
