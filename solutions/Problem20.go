// https://projecteuler.net/problem=20

package solutions

import (
	"WillStephen/GoProjectEuler/common"
	"math/big"
)

func factorial(num int) *big.Int {
	total := big.NewInt(1)
	for num > 1 {
		total = total.Mul(big.NewInt(int64(num)), total)
		num--
	}
	return total
}

// RunSolution20 finds the sum of the digits of the number 100! (100 factorial)
func RunSolution20() int {
	num := factorial(100)
	return common.SumCharacters(common.SplitBigIntToChars(num))
}
