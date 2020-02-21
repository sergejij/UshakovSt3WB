package visitor

import (
	`fmt`
	`math`

	`UshakovSt3WB/pkg/foursquare`
	"UshakovSt3WB/pkg/models"
	`UshakovSt3WB/pkg/rectangle`
	`UshakovSt3WB/pkg/rhombus`
	`UshakovSt3WB/pkg/triangle`
)

// Visitor draws geometric shapes.
type Visitor interface {
	DrawRhombusAndReturnSquare(r rhombus.Rhombus) (square int)
	DrawTriangleAndReturnSquare(t triangle.Triangle) (square float64)
	DrawRectangleAndReturnSquare(r rectangle.Rectangle) (square int)
	DrawFoursquareAndReturnSquare(f foursquare.Foursquare) (square int)
}

type visitor struct{}

/// DrawRhombus ...
func (v *visitor) DrawRhombusAndReturnSquare(r rhombus.Rhombus) (square int) {
	fmt.Println(models.NewLine, models.Rhombus)
	side := r.Side()
	square = side * side / 2
	center := side / 2
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			if i <= center {
				if j >= center-i && j <= center+i {
					fmt.Printf(models.Star)
				} else {
					fmt.Printf(models.Space)
				}
			} else {
				if (j >= center+i-side+1) && (j <= center-i+side-1) {
					fmt.Printf(models.Star)
				} else {
					fmt.Printf(models.Space)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println(models.SideRhombus, side)
	fmt.Println(models.Square, square)
	return
}

// DrawTriangle ...
func (v *visitor) DrawTriangleAndReturnSquare(t triangle.Triangle) (square float64) {
	side := t.Side()
	square = float64(side) * float64(side) * models.RootOf3 / 4
	fmt.Println(models.NewLine, models.Triangle)
	for i := 0; i < side; i += 2 {
		for j := 0; j < side; j++ {
			if (j*2 <= side && (side-i)/2 <= j) ||
				(j*2 >= side && i >= j*2-side) {
				fmt.Printf(models.Star)
			} else {
				fmt.Printf(models.Space)
			}
		}
		fmt.Println()
	}
	fmt.Println(models.SideTriangle, side)
	fmt.Printf("%s %.2f\n", models.Square, math.Round(square/0.5)*0.5)
	return
}

// DrawRectangle ...
func (v *visitor) DrawRectangleAndReturnSquare(r rectangle.Rectangle) (square int) {
	var w string
	width, height := r.Sides()
	square = width * height
	fmt.Println(models.NewLine, models.Rectangle)
	if height == 1 {
		for j := 0; j < width; j++ {
			fmt.Printf(models.Texture)
		}
		fmt.Println()
	} else {
		for i := 0; i < height; i++ {
			fmt.Printf(models.SideLine)
			if i == 0 {
				w = models.TopLine
			} else if i == height-1 {
				w = models.BottomLine
			} else {
				w = models.DoubleSpace
			}
			for j := 0; j < width; j++ {
				fmt.Printf(w)
			}
			fmt.Printf(models.SideLine)
			fmt.Println()
		}
	}
	fmt.Println(models.WidthRectangle, width)
	fmt.Println(models.HeightRectangle, height)
	fmt.Println(models.Square, square)
	return
}

// DrawFoursquare ...
func (v *visitor) DrawFoursquareAndReturnSquare(f foursquare.Foursquare) (square int) {
	var w string
	side := f.Side()
	square = side * side
	fmt.Println(models.NewLine, models.Foursquare)
	if side == 1 {
		fmt.Println(models.SmallFoursquare)
	} else {
		for i := 0; i < side; i++ {
			fmt.Printf(models.SideLine)
			if i == 0 {
				w = models.TopLine
			} else if i == side-1 {
				w = models.BottomLine
			} else {
				w = models.DoubleSpace
			}
			for j := 0; j < side; j++ {
				fmt.Printf(w)
			}
			fmt.Printf(models.SideLine)
			fmt.Println()
		}
	}
	fmt.Println(models.SideFoursquare, side)
	fmt.Println(models.Square, square)
	return
}

// NewVisitor initializes the Visitor.
func NewVisitor() Visitor {
	return &visitor{}
}
