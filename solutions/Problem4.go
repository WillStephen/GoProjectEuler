// https://projecteuler.net/problem=3

package solutions

import (
	"strconv"
)

func reverse(str string) string {
	length := len(str)
	newstr := []rune(str)
	for i, char := range []rune(str) {
		newstr[(length-i)-1] = char
	}

	return string(newstr)
}

func isPalindrome(num int) bool {
	numstr := strconv.Itoa(num)

	firstpart, secondpart := "", ""

	length := len(numstr)
	if length%2 == 0 {
		firstpart = numstr[0:(length / 2)]
		secondpart = numstr[(length / 2):length]
	} else {
		firstpart = numstr[0:(length / 2)]
		secondpart = numstr[(length+1)/2 : length]
	}

	if firstpart == reverse(secondpart) {
		return true
	}

	return false
}

func getLargestPalindrome(minFactor int, maxFactor int) int {
	maxProduct := 0

	for i := maxFactor; i >= minFactor; i-- {
		for j := maxFactor; j >= minFactor; j-- {
			product := i * j
			if isPalindrome(product) {
				if product > maxProduct {
					maxProduct = product
				}
			}
		}
	}

	return maxProduct
}

// RunSolution4 returns the largest palindromic number made from the product of
// two 3-digit numbers
func RunSolution4() int {
	return getLargestPalindrome(100, 999)
}
