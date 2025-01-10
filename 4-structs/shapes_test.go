package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 15.0}
	got := rectangle.Perimeter()
	want := 50.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := checkArea(rectangle) // use the checkArea function to calculate the area of the rectangle
		want := 72.0
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := checkArea(circle) // use the checkArea function to calculate the area of the circle
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})

	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{12, 6}
		got := checkArea(triangle)
		want := 36.0
		if got != want {
			t.Errorf("got %f, want %f", got, want)
		}
	})
}