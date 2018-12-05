package common

// Sum sums a slice of ints
func Sum(numbers []int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}

// SumRange returns the sum of all numbers between a min and max (inclusive)
func SumRange(min int, max int) int {
	total := 0
	for i := min; i <= max; i++ {
		total += i
	}

	return total
}
