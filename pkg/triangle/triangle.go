package triangle

type visitor interface {
	DrawTriangleAndReturnSquare(t Triangle) (square float64)
}

// Triangle ...
type Triangle interface {
	AcceptAndReturnSquare(v visitor) float64
	Side() int
}

type triangle struct {
	side int
}

// Accept ...
func (t *triangle) AcceptAndReturnSquare(v visitor) float64 {
	return v.DrawTriangleAndReturnSquare(t)
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
