package common

import (
	"math"
)

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

// IsPrime determines whether a number is prime
func IsPrime(num int) bool {
	if num%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
