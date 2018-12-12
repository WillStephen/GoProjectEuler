// https://projecteuler.net/problem=10

package solutions

import "WillStephen/GoProjectEuler/common"

func getPrimesBelow(limit int) []int {
	primes := []int{}
	current := 0
	for current < limit {
		if common.IsPrime(current) {
			primes = append(primes, current)
		}
		current++
	}

	return primes
}

// RunSolution10 finds the sum of all primes below 2 million
func RunSolution10() int {
	primes := getPrimesBelow(2000000)
	return common.Sum(primes)
}
