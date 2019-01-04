// https://projecteuler.net/problem=18

package solutions

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	value int
	x     int
	y     int
}

func getPointAt(points *[][]point, y int, x int) *point {
	for _, point := range (*points)[y] {
		if point.x == x {
			return &point
		}
	}
	return nil
}

func getChildren(points *[][]point, parent *point) (*point, *point) {
	if parent.y >= len(*points)-1 {
		return nil, nil
	}

	leftChild := getPointAt(points, parent.y+1, parent.x-1)
	rightChild := getPointAt(points, parent.y+1, parent.x+1)
	return leftChild, rightChild
}

func getMaximumPath(points [][]point) int {
	maxPath := 0
	for i := len(points) - 2; i >= 0; i-- {
		row := points[i]
		for j := range row {
			leftChild, rightChild := getChildren(&points, &row[j])
			if leftChild.value > rightChild.value {
				row[j].value += leftChild.value
			} else {
				row[j].value += rightChild.value
			}
			if row[j].x == 0 && row[j].y == 0 {
				maxPath = row[j].value
			}
		}
	}

	return maxPath
}

func parseTriangle(rawPoints [][]int) [][]point {
	currentid := 0
	points := [][]point{}

	for y, i := range rawPoints {
		currentRowPoints := []point{}
		for right, j := range i {
			rowlen := len(i)
			x := 0
			if rowlen%2 == 0 {
				x = right - (rowlen / 2)
				if right >= rowlen/2 {
					x++
				}
				if math.Abs(float64(x)) > 1 {
					x *= 2
					if x > 0 {
						x--
					} else {
						x++
					}
				}
			} else {
				x = right - ((rowlen - 1) / 2)
				x *= 2
			}

			currentRowPoints = append(currentRowPoints, point{j, x, y})
			currentid++
		}
		points = append(points, currentRowPoints)
	}
	return points
}

func readTriangleFromFile(filename string) [][]int {
	file, _ := os.Open(filename)
	defer file.Close()

	triangle := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []int{}
		numStrs := strings.Split(scanner.Text(), " ")
		for _, str := range numStrs {
			num, _ := strconv.Atoi(str)
			line = append(line, num)
		}
		triangle = append(triangle, line)
	}

	return triangle
}

// RunSolution18 finds the maximum path through a triangle of integers
func RunSolution18() int {
	rawTriangle := readTriangleFromFile("data/problem18_triangle.txt")

	triangle := parseTriangle(rawTriangle)
	return getMaximumPath(triangle)
}
