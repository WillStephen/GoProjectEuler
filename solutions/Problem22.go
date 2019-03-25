// https://projecteuler.net/problem=22

package solutions

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func alphabeticalValue(char string) int {
	return int([]rune(char)[0]) - 64
}

func readNamesFromFile(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return strings.Split(scanner.Text(), ",")
}

func scoreName(name string, index int) int {
	total := 0
	chars := []rune(name)
	for _, c := range chars {
		if c > 64 && c < 91 {
			total += int(c) - 64
		}
	}

	return total * (index + 1)
}

func scoreNames(names []string) int {
	total := 0
	for i, name := range names {
		total += scoreName(name, i)
	}

	return total
}

// RunSolution22 calculates a score for each name in a large file of names, and finds the sum of the scores
func RunSolution22() int {
	names := readNamesFromFile("data/problem22_names.txt")

	sort.Strings(names)
	return scoreNames(names)
}
