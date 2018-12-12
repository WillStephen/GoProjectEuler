// https://projecteuler.net/problem=1

package solutions

import (
	"WillStephen/GoProjectEuler/common"
)

func getMultiplesOf(upperLim int, targets ...int) []int {
	multiples := []int{}

	for i := 1; i <= upperLim; i++ {
		for _, target := range targets {
			if i%target == 0 {
				multiples = append(multiples, i)
				break
			}
		}
	}

	return multiples
}

// RunSolution1 returns the sum of all multiples of 3 and 5 below 1000
func RunSolution1() int {
	multiples := getMultiplesOf(999, 3, 5)
	return common.Sum(multiples)
}
