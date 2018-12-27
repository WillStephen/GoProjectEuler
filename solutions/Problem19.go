// https://projecteuler.net/problem=19

package solutions

var months = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

var days = map[int]string{
	0: "monday",
	1: "tuesday",
	2: "wednesday",
	3: "thursday",
	4: "friday",
	5: "saturday",
	6: "sunday",
}

func isLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	return year%4 == 0
}

func countDays(dayOfWeek int, startYear int, endYear int) int {
	total := 0
	currentDayOfWeek := 0
	for year := startYear; year <= endYear; year++ {
		for currentMonth := 1; currentMonth <= 12; currentMonth++ {
			if isLeapYear(year) {
				months[2] = 29
			} else {
				months[2] = 28
			}
			for currentDayOfMonth := 1; currentDayOfMonth <= months[currentMonth]; currentDayOfMonth++ {
				currentDayOfWeek++
				if currentDayOfWeek == dayOfWeek && currentDayOfMonth == 1 {
					total++
				}
				if currentDayOfWeek > 6 {
					currentDayOfWeek = 0
				}
			}
		}
	}

	return total
}

// RunSolution19 finds the number of Sundays that fell on the first of the month in the twentieth century
func RunSolution19() int {
	return countDays(6, 1901, 2000)
}
