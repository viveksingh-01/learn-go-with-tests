package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := []string{"Chris", "London"}
	x := struct {
		Name string
		City string
	}{"Chris", "London"}

	var got []string
	callbackFn := func(input string) {
		got = append(got, input)
	}

	Walk(x, callbackFn)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, want %v", got, expected)
	}
}
