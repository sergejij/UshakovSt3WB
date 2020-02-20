package rhombus

type visitor interface {
	DrawRhombus(r Rhombus)
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
	v.DrawRhombus(r)
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
