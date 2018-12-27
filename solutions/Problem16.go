// https://projecteuler.net/problem=16

package solutions

import (
	"WillStephen/GoProjectEuler/common"
	"math/big"
)

func getPower(base int, power int) *big.Int {
	bigBase := big.NewInt(int64(base))
	bigPower := big.NewInt(int64(power))
	return new(big.Int).Exp(bigBase, bigPower, nil)
}

// RunSolution16 finds sum of the digits of 2^1000
func RunSolution16() int {
	num := getPower(2, 1000)
	digitChars := common.SplitBigIntToChars(num)
	return common.SumCharacters(digitChars)
}
