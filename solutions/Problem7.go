// https://projecteuler.net/problem=7

package solutions

import (
	"GoProjectEuler/common"
)

func getNthPrime(n int) int {
	nthPrimeFound := false
	primeCount := 1 // 2 is the only even prime
	currentNumber := 1

	for !nthPrimeFound {
		if common.IsPrime(currentNumber) {
			primeCount++
		}

		if primeCount == n {
			nthPrimeFound = true
		} else {
			currentNumber += 2
		}
	}

	return currentNumber
}

// RunSolution7 finds the 10001st prime number
func RunSolution7() int {
	return getNthPrime(10001)
}
