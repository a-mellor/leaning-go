package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// Return the perimiter of a given width and height
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Return the area of a given width and height
func (rectangle *Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func Area(circle Circle) float64 {
	return math.Pi * circle.Radius * circle.Radius
}
