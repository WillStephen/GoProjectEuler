// https://projecteuler.net/problem=5

package solutions

func isDivisibleByAll(num int, min int, max int) bool {
	if num < max {
		return false
	}

	for i := min; i <= max; i++ {
		if num%i != 0 {
			return false
		}
	}

	return true
}

func getSmallestNumberDivisibleByRange(min int, max int) int {
	num := max * max
	for true {
		if isDivisibleByAll(num, 1, 20) {
			return num
		}

		num += max
	}

	return 0
}

// RunSolution5 returns the smallest number divisible by all numbers 1-20 inclusive
func RunSolution5() int {
	return getSmallestNumberDivisibleByRange(1, 20)
}
