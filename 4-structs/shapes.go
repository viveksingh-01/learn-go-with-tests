package structs

import "math"

// A struct is just a named collection of fields where you can store data.
// They are a way to create more complex data types.
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {	
	Base   float64
	Height float64
}

// Instead of defining 2 different methods to calculate Area,
// we want to be able to write some kind of checkArea function that we can pass 
// both Rectangles and Circles to, but fail to compile if we try to pass in something that isn't a shape.

// Interfaces are very powerful concept in statically-typed languages like Go.
// They allow us to make functions that can be used with different types and create highly-decoupled code.

// But how does something become a shape? We just tell Go what a Shape is, using an interface declaration
// We can do this by defining a Shape interface that defines a method Area() that returns a float64.
// Any type that defines a method Area() that returns a float64 satisfies the Shape interface.
type Shape interface {
	Area() float64
}

// Now we can define a function checkArea that takes a Shape and returns a float64.
// We can pass in any type that satisfies the Shape interface, like Rectangle or Circle.
// If we try to pass in something that doesn't satisfy the Shape interface, the Go compiler will throw an error.
// This is a great way to make sure our code is correct and decoupled.
func checkArea(s Shape) float64 {
	return s.Area()
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// More on Interfaces:
// This is a powerful concept in Go and is used throughout the standard library.
// For example, the fmt package uses interfaces to print different types of data with the same function.
// This is why fmt.Println can print a string, an integer, a float, etc.
// The io package uses interfaces to read and write data from different sources like files, network connections, and in-memory data.