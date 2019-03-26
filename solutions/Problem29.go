// https://projecteuler.net/problem=29

package solutions

import "math/big"

func removeDuplicates(numbers []string) []string {
	numberMap := make(map[string]bool)
	uniques := []string{}
	for _, num := range numbers {
		if _, exists := numberMap[num]; !exists {
			uniques = append(uniques, num)
			numberMap[num] = true
		}
	}

	return uniques
}

func getTerms(aFloor int, aCeiling int, bFloor int, bCeiling int) []string {
	terms := []string{}

	for a := aFloor; a <= aCeiling; a++ {
		for b := bFloor; b <= bCeiling; b++ {
			bigA := big.NewInt(int64(a))
			bigB := big.NewInt(int64(b))
			result := bigA.Exp(bigA, bigB, nil)
			terms = append(terms, result.String())
		}
	}

	return terms
}

// RunSolution29 finds the number of distinct terms in the sequence a^b for 2 <= a and b <= 100
func RunSolution29() int {
	terms := getTerms(2, 100, 2, 100)
	return len(removeDuplicates(terms))
}
