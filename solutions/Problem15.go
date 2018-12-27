// https://projecteuler.net/problem=15

package solutions

import (
	"math"
)

func getPaths(gridSteps int, currentSteps int, currentScore int, count int) int {
	if currentSteps == gridSteps {
		if int(math.Abs(float64(currentScore))) > (gridSteps - currentSteps) {
			return count
		}

		if currentScore == 0 {
			if count%50 == 0 {
				println(count)
			}
			count = count + 1
		}
		return count
	}

	rightScore := currentScore + 1
	count = getPaths(gridSteps, currentSteps+1, rightScore, count)

	downScore := currentScore - 1
	count = getPaths(gridSteps, currentSteps+1, downScore, count)

	return count
}

func getNumberOfPathsThroughGrid(gridSize int) int {
	//return len(getPaths(gridSize*2, 0, []string{}, 0, [][]string{{}})[1:])

	return getPaths(gridSize*2, 0, 0, 0)
}

// RunSolution15 finds the number of possible paths through a 20x20 grid, moving only right and down
func RunSolution15() int {
	//return getPaths(6, 0, []string{}, 0, 0)
	return getNumberOfPathsThroughGrid(20)
}
