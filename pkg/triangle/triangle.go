package triangle

type visitor interface {
	DrawTriangle(t Triangle)
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
	v.DrawTriangle(t)
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
