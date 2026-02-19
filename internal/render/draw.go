package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/piotrowski/ebitris/internal/tetris"
)

const BlockSize = 30

var borderColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}

func DrawBlock(screen *ebiten.Image, x, y int, rectangleColor color.Color) {
	pixelX := float32(x * BlockSize)
	pixelY := float32(y * BlockSize)

	vector.FillRect(screen, pixelX, pixelY, BlockSize, BlockSize, rectangleColor, true)
	vector.StrokeRect(screen, pixelX, pixelY, BlockSize, BlockSize, 2, borderColor, true)
}

func DrawBoard(screen *ebiten.Image, board *tetris.Board, offsetX, offsetY int) {
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			cellValue := board.Cell(x, y)

			bgColor := color.RGBA{R: 20, G: 20, B: 30, A: 255}
			DrawBlock(screen, offsetX+x, offsetY+y, bgColor)

			if cellValue != 0 {
				pieceColor := tetris.GetPieceColor(cellValue)
				DrawBlock(screen, offsetX+x, offsetY+y, pieceColor)
			}
		}
	}
}

func DrawPiece(screen *ebiten.Image, piece *tetris.Piece, offsetX, offsetY int) {
	colorPiece := tetris.GetPieceColor(piece.Color)

	for _, cell := range piece.GetCells() {
		x := piece.X + cell.X + offsetX
		y := piece.Y + cell.Y + offsetY

		if y >= 0 {
			DrawBlock(screen, x, y, colorPiece)
		}
	}
}

func DrawText(screen *ebiten.Image, textToDraw string, x, y int, fontFace text.Face) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))

	text.Draw(screen, textToDraw, fontFace, op)
}
