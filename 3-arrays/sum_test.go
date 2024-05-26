package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of collection of 5 numbers", func(t *testing.T) {
		// Arrays have a fixed capacity which you define when you declare the variable. 
		// We can initialize an array in two ways:
		// [N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
		// [...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}
		numArray := [5]int{1, 2, 3, 4,5 }
		
		got := SumOfArray(numArray)
		want := 15
		if got != want {
			// It is sometimes useful to also print the inputs to the function in the error message. 
			// Here, we are using the %v placeholder to print the "default" format, which works well for arrays.
			t.Errorf("got %d want %d, given %v", got, want, numArray)
		}
	})

	// NOTE:
	// An interesting property of arrays is that the size is encoded in its type. 
	// If you try to pass an [4]int into a function that expects [5]int, it won't compile. 
	// They are different types so it's just the same as trying to pass a string into a function that wants an int.

	t.Run("sum of collection of any size", func(t *testing.T) {
		// Go has "slices" which do not encode the size of the collection and instead can have any size.
		numSlice := []int{1, 2, 3}
		got := SumOfSlice(numSlice)
		want := 6
		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numSlice)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// "slices.Equal" function can be used to do a simple shallow compare on slices
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}