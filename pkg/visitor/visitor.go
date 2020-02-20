package visitor

import (
	`fmt`

	`visitor/pkg/foursquare`
	"visitor/pkg/models"
	`visitor/pkg/rectangle`
	`visitor/pkg/rhombus`
	`visitor/pkg/triangle`
)

// Visitor draws geometric shapes.
type Visitor interface {
	DrawRhombus(r rhombus.Rhombus)
	DrawTriangle(t triangle.Triangle)
	DrawRectangle(r rectangle.Rectangle)
	DrawFoursquare(f foursquare.Foursquare)
}

type visitor struct {
}

/// DrawRhombus ...
func (v *visitor) DrawRhombus(r rhombus.Rhombus) {
	fmt.Println(models.NewLine, models.Rhombus)
	side := r.Side()
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
	fmt.Println(models.SideRhombus, side*side/2)
}

// DrawTriangle ...
func (v *visitor) DrawTriangle(t triangle.Triangle) {
	side := t.Side()
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
	fmt.Printf("%s %.2f\n", models.Square, (float64(side)*float64(side)*models.RootOf3)/4)
}

// DrawRectangle ...
func (v *visitor) DrawRectangle(r rectangle.Rectangle) {
	var w string
	width, height := r.Sides()
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
	fmt.Println(models.Square, width*height)
}

// DrawFoursquare ...
func (v *visitor) DrawFoursquare(f foursquare.Foursquare) {
	var w string
	side := f.Side()
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
	fmt.Println(models.Square, side*side)
}

// NewVisitor initializes the Visitor.
func NewVisitor() Visitor {
	return &visitor{}
}
