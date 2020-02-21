package visitor

import (
	`math`
	`testing`

	`github.com/stretchr/testify/assert`

	`UshakovSt3WB/pkg/foursquare`
	`UshakovSt3WB/pkg/rectangle`
	`UshakovSt3WB/pkg/rhombus`
	`UshakovSt3WB/pkg/triangle`
)

const (
	testOfRhombus    = "Test Rhombus"
	testOfTriangle   = "Test Triangle"
	testOfFoursquare = "Test Foursquare"
	testOfRectangle  = "Test Rectangle"
)

var (
	parametersOfRhombus = map[string]int{
		"side":       10,
		"wantSquare": 50,
	}
	parametersOfTriangle = map[string]float64{
		"side":       10,
		"wantSquare": 43.5,
	}
	parametersOfFoursquare = map[string]int{
		"side":       10,
		"wantSquare": 100,
	}
	parametersOfRectangle = map[string]int{
		"height":     5,
		"width":      15,
		"wantSquare": 75,
	}
)

func TestVisitor(t *testing.T) {
	visitor := NewVisitor()
	t.Run(testOfRhombus, func(t *testing.T) {
		newRhombus := rhombus.NewRhombus(parametersOfRhombus["side"])
		square := newRhombus.AcceptAndReturnSquare(visitor)
		assert.Equal(t, parametersOfRhombus["wantSquare"], square)
	})
	t.Run(testOfTriangle, func(t *testing.T) {
		newTriangle := triangle.NewTriangle(int(parametersOfTriangle["side"]))
		square := newTriangle.AcceptAndReturnSquare(visitor)
		assert.Equal(t, parametersOfTriangle["wantSquare"], math.Round(square/0.5)*0.5)
	})
	t.Run(testOfFoursquare, func(t *testing.T) {
		newFoursquare := foursquare.NewFoursquare(parametersOfFoursquare["side"])
		square := newFoursquare.AcceptAndReturnSquare(visitor)
		assert.Equal(t, parametersOfFoursquare["wantSquare"], square)
	})
	t.Run(testOfRectangle, func(t *testing.T) {
		rhombus := rectangle.NewRectangle(parametersOfRectangle["height"],
			parametersOfRectangle["width"])
		square := rhombus.AcceptAndReturnSquare(visitor)
		assert.Equal(t, parametersOfRectangle["wantSquare"], square)
	})
}
