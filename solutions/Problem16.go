// https://projecteuler.net/problem=16

package solutions

import (
	"math/big"
	"strconv"
	"strings"
)

func getPower(base int, power int) *big.Int {
	bigBase := big.NewInt(int64(base))
	bigPower := big.NewInt(int64(power))
	return new(big.Int).Exp(bigBase, bigPower, nil)
}

func splitBigIntToChars(num *big.Int) []string {
	return strings.Split(num.String(), "")
}

func sumCharacters(numChars []string) int {
	total := 0
	for _, i := range numChars {
		num, _ := strconv.Atoi(i)
		total += num
	}

	return total
}

// RunSolution16 finds sum of the digits of 2^1000
func RunSolution16() int {
	num := getPower(2, 1000)
	digitChars := splitBigIntToChars(num)
	return sumCharacters(digitChars)
}
