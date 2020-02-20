package rectangle

type visitor interface {
	DrawRectangle(r Rectangle)
}

// Rectangle ...
type Rectangle interface {
	Accept(v visitor)
	Sides() (int, int)
}

type rectangle struct {
	height int
	width  int
}

// Accept ...
func (r *rectangle) Accept(v visitor) {
	v.DrawRectangle(r)
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
