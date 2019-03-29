// https://projecteuler.net/problem=30

package solutions

import (
	"WillStephen/GoProjectEuler/common"
	"math"
	"strconv"
	"strings"
)

func findMaxNumberToCheck(power int) int {
	found := false
	numOfDigits := 1
	base := 9
	for !found {
		testNumber, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(base), numOfDigits))
		result := int(math.Pow(float64(base), float64(power))) * numOfDigits
		if result < testNumber {
			return testNumber
		}
		numOfDigits++
	}

	return -1
}

func isSumOfDigitsToNthPower(num int, power int) bool {
	str := strconv.Itoa(num)
	total := 0
	for _, char := range str {
		digit, _ := strconv.Atoi(string(char))
		total += int(math.Pow(float64(digit), float64(power)))
	}

	return total == num
}

// RunSolution30 returns the sum of all numbers that can be written as the sum of fifth powers of their digits
func RunSolution30() int {
	maxNum := findMaxNumberToCheck(5)
	nums := []int{}
	for i := 2; i <= maxNum; i++ {
		if isSumOfDigitsToNthPower(i, 5) {
			nums = append(nums, i)
		}
	}

	return common.Sum(nums)
}
