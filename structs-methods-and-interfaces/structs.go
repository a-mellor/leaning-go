package structs

import "math"

// Shape is implemented by anything that can tell us its Area
type Shape interface {
	Area() float64
}

// Rectangle has the dimensions of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a circle
type Circle struct {
	Radius float64
}

// Returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle has the dimensions of a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Returns the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// Returns the perimiter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
