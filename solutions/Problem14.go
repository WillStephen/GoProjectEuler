// https://projecteuler.net/problem=11

package solutions

func getNextCollatzNumber(num int) int {
	if num%2 == 0 {
		return num / 2
	}
	return (num * 3) + 1
}

func getCollatzSequence(starting int) []int {
	sequence := []int{starting}
	current := starting

	for current != 1 {
		current = getNextCollatzNumber(current)
		sequence = append(sequence, current)
	}

	return sequence
}

func getNumberWithLongestCollatzSequence(max int) int {
	numberWithLongestCollatzSequence := 1
	longestSequenceLength := len(getCollatzSequence(1))

	for i := 2; i < max; i++ {
		length := len(getCollatzSequence(i))
		if length > longestSequenceLength {
			numberWithLongestCollatzSequence = i
			longestSequenceLength = length
		}
	}

	return numberWithLongestCollatzSequence
}

// RunSolution14 finds the number (under one million) that produces the longest Collatz sequence
func RunSolution14() int {
	return getNumberWithLongestCollatzSequence(1000000)
}
