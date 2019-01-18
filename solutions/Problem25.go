// https://projecteuler.net/problem=25

package solutions

import "math/big"

func getDigitCount(num *big.Int) int {
	return len(num.String())
}

func getIndexOfFirstFibonacciTermWithDigits(digits int) int {
	num1 := big.NewInt(1)
	num2 := big.NewInt(1)

	index := 2
	done := false
	for !done {
		num3 := num1.Add(num1, num2)
		index++
		if getDigitCount(num3) == digits {
			done = true
		}

		num1 = num2
		num2 = num3
	}
	return index
}

// RunSolution25 finds the index of the first 1000-digit Fibonacci number
func RunSolution25() int {
	return getIndexOfFirstFibonacciTermWithDigits(1000)
}
