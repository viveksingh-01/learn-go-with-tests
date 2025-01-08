package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 15.0}
	got := Perimeter(rectangle)
	want := 50.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := Area(rectangle)
		want := 72.0
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		// got := Area(circle) --> this will throw error
		// since we cannot use circle (variable of type Circle) as Rectangle value in argument to Area
		// also, we can't declare another Area function with Circle as argument.
		// so, we have only two options:
		// 1. define another Area function with Circle as argument in a different package
		// 2. use methods (Receiver functions) to define Area function with Circle as argument.
		// we will use the second option.
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}