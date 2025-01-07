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
	rectangle := Rectangle{10.0, 15.0}
	got := Area(rectangle)
	want := 150.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}