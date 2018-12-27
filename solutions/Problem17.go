// https://projecteuler.net/problem=17

package solutions

import (
	"strconv"
	"strings"
)

var connector = 3
var thousand = 8
var hundred = 7

var numMap = map[string][]int{
	"1": []int{3, 6, 0},
	"2": []int{3, 6, 6},
	"3": []int{5, 8, 6},
	"4": []int{4, 8, 5},
	"5": []int{4, 7, 5},
	"6": []int{3, 7, 5},
	"7": []int{5, 9, 7},
	"8": []int{5, 8, 6},
	"9": []int{4, 8, 6},
	"0": []int{0, 3, 0},
}

func splitToChars(num int) []string {
	str := strconv.Itoa(num)
	return strings.Split(str, "")
}

func countCharsInNumber(num int) int {
	chars := splitToChars(num)

	total := 0
	for i, char := range chars {
		index := len(chars) - i
		if index == 1 {
			if len(chars) > 1 && chars[len(chars)-2] == "1" {
				total += numMap[char][1]
			} else {
				total += numMap[char][0]
			}
			continue
		}
		if index == 2 {
			total += numMap[char][2]
		} else {
			total += numMap[char][0]
			if index == 3 && char != "0" {
				total += hundred
				if (chars[len(chars)-1] != "0") || (chars[len(chars)-2] != "0") {
					total += connector
				}
			}
			if index == 4 {
				total += thousand
			}
		}
	}

	return total
}

func sumCharsInRange(min int, max int) int {
	total := 0
	for i := min; i <= max; i++ {
		total += countCharsInNumber(i)
	}
	return total
}

// RunSolution17 finds the number of letters used when writing the numbers 1-1000 in words
func RunSolution17() int {
	print(sumCharsInRange(1, 1000))
	return sumCharsInRange(1, 1000)
}
