// https://projecteuler.net/problem=11

package solutions

import "math"

func getFactors(num int) []int {
	factors := []int{}
	sqrt := math.Sqrt(float64(num))

	for i := 1; i <= int(math.Floor(sqrt)); i++ {
		if num%i == 0 {
			factors = append(factors, i, num/i)
		}
	}

	if sqrt*sqrt == float64(num) {
		factors = append(factors, int(sqrt))
	}

	return factors
}

func getFirstTriangleNumberWithMoreDivisorsThan(divisors int) int {
	counter := 1
	found := false
	triangularNum := 0

	for !found {
		triangularNum = triangularNum + counter
		factors := getFactors(triangularNum)
		if len(factors) > divisors {
			break
		}
		counter++
	}

	return triangularNum
}

// RunSolution12 finds the smallest triangle number with over 500 divisors
func RunSolution12() int {
	return getFirstTriangleNumberWithMoreDivisorsThan(500)
}
