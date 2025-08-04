package shapes

import "math"

// Shape is an interface for all the methods to implement the
// Area method. This is the GO method of setting interfaces
type Shape interface {
	Area() float64
}

// Types
type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Triangle struct {
	Base   float64
	Height float64
}

// Methods
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * rectangle.Width * rectangle.Height
}
