package common

import (
	"math"
	"math/big"
	"strconv"
	"strings"
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
	// special cases
	if num == 1 {
		return false
	}
	if num == 2 {
		return true
	}

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

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// Min returns the smallest integer
func Min(nums ...int) int {
	min := nums[0]
	for _, i := range nums {
		if i < min {
			min = i
		}
	}
	return min
}

// Max returns the largest integer
func Max(nums ...int) int {
	max := nums[0]
	for _, i := range nums {
		if i > max {
			max = i
		}
	}
	return max
}

// SplitBigIntToChars splits a big int into a slice of strings, each of length 1
func SplitBigIntToChars(num *big.Int) []string {
	return strings.Split(num.String(), "")
}

// SumCharacters sums a slice of number strings
func SumCharacters(numChars []string) int {
	total := 0
	for _, i := range numChars {
		num, _ := strconv.Atoi(i)
		total += num
	}

	return total
}

// Factorial returns the factorial of a number
func Factorial(num int) int {
	for i := num - 1; i > 1; i-- {
		num *= i
	}

	return num
}
