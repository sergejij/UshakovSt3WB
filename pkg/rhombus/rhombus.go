package rhombus

type visitor interface {
	DrawRhombusAndReturnSquare(r Rhombus) (square int)
}

// Rhombus ...
type Rhombus interface {
	Accept(v visitor)
	Side() int
}

type rhombus struct {
	side int
}

// Accept ...
func (r *rhombus) Accept(v visitor) {
	v.DrawRhombusAndReturnSquare(r)
}

// Side ...
func (r *rhombus) Side() int {
	return r.side
}

// NewRhombus initializes the Rhombus.
func NewRhombus(side int) Rhombus {
	return &rhombus{
		side: side,
	}
}
