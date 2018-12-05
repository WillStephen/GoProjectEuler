// https://projecteuler.net/problem=6

package solutions

import (
	"GoProjectEuler/common"
)

func getSquares(first int, last int) []int {
	var squares []int
	for i := first; i <= last; i++ {
		squares = append(squares, i*i)
	}

	return squares
}

// RunSolution6 finds the difference between the sum of the squares and the
// square of the sum of the natural numbers 1-100
func RunSolution6() int {
	squares := getSquares(1, 100)
	sumSquares := common.Sum(squares)

	sum := common.SumRange(1, 100)
	squareSum := sum * sum

	return squareSum - sumSquares
}
