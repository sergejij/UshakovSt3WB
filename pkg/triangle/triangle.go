package triangle

type visitor interface {
	DrawTriangleAndReturnSquare(t Triangle) (square float64)
}

// Triangle ...
type Triangle interface {
	Accept(v visitor)
	Side() int
}

type triangle struct {
	side int
}

// Accept ...
func (t *triangle) Accept(v visitor) {
	v.DrawTriangleAndReturnSquare(t)
}

// Side ...
func (t *triangle) Side() int {
	return t.side
}

// NewTriangle initializes the Triangle.
func NewTriangle(side int) Triangle {
	return &triangle{
		side: side,
	}
}
