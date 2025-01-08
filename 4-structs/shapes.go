package structs

import "math"

// A struct is just a named collection of fields where you can store data.
// They are a way to create more complex data types.
type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

type Circle struct {
	Radius float64
}

// A method is a function with a special receiver argument.
// Methods are very similar to functions but they are called by invoking them on an instance of a particular type.
// Here we are defining the 'Area' method with receiver type 'Circle'.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}