package main

import (
	`visitor/pkg/foursquare`
	`visitor/pkg/rectangle`
	`visitor/pkg/rhombus`
	`visitor/pkg/triangle`

	`visitor/pkg/visitor`
)

const (
	sideOfRhombus     = 15
	sideOfTriangle    = 24
	sideOfFoursquare  = 9
	heightOfRectangle = 6
	widthOfRectangle  = 10
)

func main() {
	newVisitor := visitor.NewVisitor()
	newRhombus := rhombus.NewRhombus(sideOfRhombus)
	newRectangle := rectangle.NewRectangle(heightOfRectangle, widthOfRectangle)
	newTriangle := triangle.NewTriangle(sideOfTriangle)
	newFoursquare := foursquare.NewFoursquare(sideOfFoursquare)

	newTriangle.Accept(newVisitor)
	newRectangle.Accept(newVisitor)
	newFoursquare.Accept(newVisitor)
	newRhombus.Accept(newVisitor)
}
