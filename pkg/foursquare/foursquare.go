package foursquare

type visitor interface {
	DrawFoursquareAndReturnSquare(f Foursquare) (square int)
}

// Foursquare ...
type Foursquare interface {
	Accept(v visitor)
	Side() int
}

type foursquare struct {
	side int
}

// Accept ...
func (f *foursquare) Accept(v visitor) {
	v.DrawFoursquareAndReturnSquare(f)
}

// Side ...
func (f *foursquare) Side() int {
	return f.side
}

// NewFoursquare initializes the Foursquare.
func NewFoursquare(side int) Foursquare {
	return &foursquare{
		side: side,
	}
}
