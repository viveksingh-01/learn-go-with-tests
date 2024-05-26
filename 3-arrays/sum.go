package arrays

func SumOfArray(numbers [5]int) int {
	sum := 0
	// "range" lets you iterate over an array.
	// On each iteration, range returns two values - the index and the value.
	// We are choosing to ignore the index value by using "_" blank identifier.
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumOfSlice(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Go lets us write variadic functions that can take a variable number of arguments.
func SumAll(numbersToSum ...[]int) []int {
	// "make" allows you to create a slice with a starting capacity of the length we need to work through.
	// sums := make([]int, len(numbersToSum))

	// NOTE: As mentioned, slices have a capacity.
	// If we have a slice with a capacity of 2 and try to do mySlice[10] = 1 we will get a runtime error.
	// However, we can use the "append" function which takes a slice and a new value, and
	// then returns a new slice with all the items in it.

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, SumOfSlice((numbers)))
	}
	return sums
}