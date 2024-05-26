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
	lengthOfSlice := len(numbersToSum)
	// "make" allows you to create a slice with a starting capacity of the length we need to work through.
	sums := make([]int, lengthOfSlice)

	for i, numbers := range numbersToSum {
		sums[i] = SumOfSlice((numbers))
	}
	return sums
}