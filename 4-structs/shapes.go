package structs

import "math"

// A struct is just a named collection of fields where you can store data.
// They are a way to create more complex data types.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// A method is a function with a special receiver argument.
// Methods are very similar to functions but they are called by invoking them on an instance of a particular type.
// Here we are defining the 'Area' method with receiver type 'Circle'.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}