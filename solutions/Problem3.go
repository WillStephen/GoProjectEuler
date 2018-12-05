// https://projecteuler.net/problem=3

package solutions

func isPrime(num int) bool {
	for i := 2; i < (num / 2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func getLargestPrimeFactor(num int) int {
	for i := 2; i < (num / 2); i++ {

		if num%i != 0 {
			continue
		}

		factor := num / i
		if isPrime(factor) {
			return factor
		}
	}

	return 0
}

// RunSolution3 returns largest prime factor of 600851475143
func RunSolution3() int {
	const target = 600851475143
	return getLargestPrimeFactor(target)
}
