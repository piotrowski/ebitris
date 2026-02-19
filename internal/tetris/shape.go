package tetris

type ShapeType int

const (
	ShapeI ShapeType = iota // Straight line
	ShapeO                  // Square
	ShapeT                  // T-shape
	ShapeS                  // S-shape
	ShapeZ                  // Z-shape
	ShapeJ                  // J-shape
	ShapeL                  // L-shape
)

var shapes = map[ShapeType][][]Cell{
	ShapeI: {
		// Rotation 0: Horizontal
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}},
		// Rotation 1: Vertical
		{{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}, {X: 2, Y: 3}},
	},
	ShapeO: {
		// Only 1 rotation (square)
		{{X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 2}},
	},
	ShapeT: {
		// Rotation 0: T pointing up
		{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		// Rotation 1: T pointing right
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 2}},
		// Rotation 2: T pointing down
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 2}},
		// Rotation 3: T pointing left
		{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}},
	},
	ShapeS: {
		// Rotation 0: S horizontal
		{{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
		// Rotation 1: S vertical
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 2, Y: 2}},
	},
	ShapeZ: {
		// Rotation 0: Z horizontal
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		// Rotation 1: Z vertical
		{{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}},
	},
	ShapeJ: {
		// Rotation 0: J pointing up
		{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		// Rotation 1: J pointing right
		{{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}},
		// Rotation 2: J pointing down
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 2, Y: 2}},
		// Rotation 3: J pointing left
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 2}},
	},
	ShapeL: {
		// Rotation 0: L pointing up
		{{X: 2, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		// Rotation 1: L pointing right
		{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 2}},
		// Rotation 2: L pointing down
		{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 0, Y: 2}},
		// Rotation 3: L pointing left
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}},
	},
}
