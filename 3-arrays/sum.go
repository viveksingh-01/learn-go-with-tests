package arrays

func Sum(numbers [5]int) int {
	sum := 0
	// "range" lets you iterate over an array.
	// On each iteration, range returns two values - the index and the value.
	// We are choosing to ignore the index value by using "_" blank identifier.
	for _, num := range numbers {
		sum += num
	}
	return sum
}