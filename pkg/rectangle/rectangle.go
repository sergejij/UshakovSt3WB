package rectangle

type visitor interface {
	DrawRectangleAndReturnSquare(r Rectangle) (square int)
}

// Rectangle ...
type Rectangle interface {
	AcceptAndReturnSquare(v visitor) int
	Sides() (int, int)
}

type rectangle struct {
	height int
	width  int
}

// Accept ...
func (r *rectangle) AcceptAndReturnSquare(v visitor) int {
	return v.DrawRectangleAndReturnSquare(r)
}

// Sides ...
func (r *rectangle) Sides() (int, int) {
	return r.width, r.height
}

// NewRectangle initializes the Rectangle.
func NewRectangle(height, width int) Rectangle {
	return &rectangle{
		height: height,
		width:  width,
	}
}
