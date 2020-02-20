package shapes

type visitor interface {
	DrawRhombus(side int)
	DrawTriangle(side int)
	DrawRectangle(height, width  int)
	DrawFoursquare(side int)
}

type foursquare interface {
	Accept(v visitor)
}
type triangle interface {
	Accept(v visitor)
}
type rhombus interface {
	Accept(v visitor)
}
type rectangle struct {
	height int
	width  int
}

// Shapes ...
type element interface {
	Accept(v visitor) string
}
type Shapes struct {
	romb rhombus
	square foursquare
	triu triangle
	rect rectangle
}

func (s *Shapes) Accept(v visitor) {
	elements := []element{
		&s.rect,

	}
}
