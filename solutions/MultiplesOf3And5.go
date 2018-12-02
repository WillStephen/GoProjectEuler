// https://projecteuler.net/problem=1

package main

import (
	"fmt"
)

// getMultiplesOf takes an upper limit and any number of target numbers, and
// finds any number between 1 and the upper limit that are multiples of any
// of the target numbers
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

// sum returns the sum of a slice of numbers
func sum(numbers []int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}

func main() {
	multiples := getMultiplesOf(999, 3, 5)
	sum := sum(multiples)

	fmt.Println(sum)
}
