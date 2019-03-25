// https://projecteuler.net/problem=28

package solutions

import (
	"math"
)

func getSpiral(size int) int {
	currentNumber := 1

	currentLineLength := 1
	currentPos := 0
	xDirection, yDirection := 1, 0
	x, y := 0, 0

	xflag, yflag := -1, -1

	sumOfDiagonals := 0
	for currentLineLength <= size && x <= (size-1)/2 {

		if math.Abs(float64(x)) == math.Abs(float64(y)) {
			sumOfDiagonals += currentNumber
		}

		if currentPos == currentLineLength {
			if xDirection == 0 {
				xDirection += xflag
				xflag *= -1

				yDirection = 0
				currentLineLength++
			} else {
				yDirection += yflag
				yflag *= -1
				xDirection = 0
			}

			currentPos = 0
		}

		x += xDirection
		y += yDirection
		currentPos++
		currentNumber++
	}

	return sumOfDiagonals
}

// RunSolution28 finds the sum of the numbers on the diagonals of a 1001x1001 clockwise spiral of numbers
func RunSolution28() int {
	return getSpiral(1001)
}
