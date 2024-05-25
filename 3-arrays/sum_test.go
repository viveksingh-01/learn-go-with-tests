package arrays

import "testing"

func TestSum(t *testing.T) {
	// Arrays have a fixed capacity which you define when you declare the variable. 
	// We can initialize an array in two ways:
	// [N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
	// [...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}
	numbers := [5]int{1, 2, 3, 4,5 }

	got := Sum(numbers)
	want := 15
	if got != want {
		// It is sometimes useful to also print the inputs to the function in the error message. Here, we are using the %v placeholder to print the "default" format, which works well for arrays.
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}