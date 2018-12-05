// https://projecteuler.net/problem=2

package solutions

import (
	"GoProjectEuler/common"
)

func isEven(num int) bool {
	return num%2 == 0
}

func getEvenFibonacciNumbers(num1 int, num2 int, limit int) []int {
	var fibonacciNumbers []int

	num3 := 0
	for num3 < limit {
		num3 = num1 + num2
		if isEven(num3) {
			fibonacciNumbers = append(fibonacciNumbers, num3)
		}

		num1 = num2
		num2 = num3
	}

	return fibonacciNumbers
}

// RunSolution2 returns the sum of all even Fibonacci numbers below 4 million
func RunSolution2() int {
	evenFibonacciNumbers := getEvenFibonacciNumbers(1, 1, 4000000)
	return common.Sum(evenFibonacciNumbers)
}
