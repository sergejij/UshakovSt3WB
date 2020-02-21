package main

import (
	`UshakovSt3WB/pkg/foursquare`
	`UshakovSt3WB/pkg/rectangle`
	`UshakovSt3WB/pkg/rhombus`
	`UshakovSt3WB/pkg/triangle`

	`UshakovSt3WB/pkg/visitor`
)

const (
	sideOfRhombus     = 10
	sideOfTriangle    = 10
	sideOfFoursquare  = 10
	heightOfRectangle = 5
	widthOfRectangle  = 15
)

func main() {
	newVisitor := visitor.NewVisitor()
	newRhombus := rhombus.NewRhombus(sideOfRhombus)
	newRectangle := rectangle.NewRectangle(heightOfRectangle, widthOfRectangle)
	newTriangle := triangle.NewTriangle(sideOfTriangle)
	newFoursquare := foursquare.NewFoursquare(sideOfFoursquare)

	newTriangle.AcceptAndReturnSquare(newVisitor)
	newRectangle.AcceptAndReturnSquare(newVisitor)
	newFoursquare.AcceptAndReturnSquare(newVisitor)
	newRhombus.AcceptAndReturnSquare(newVisitor)
}
