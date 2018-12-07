// https://projecteuler.net/problem=9

package solutions

func isPythagoreanTriplet(a int, b int, c int) bool {
	return (a*a)+(b*b) == (c * c)
}

func getProductOfPythagoreanTripletWithSum(sum int) int {
	found := false

	a := 1
	b := 1
	c := 3
	for !found {
		for b = 1; b < c; b++ {
			for a = 1; a < b; a++ {
				if !isPythagoreanTriplet(a, b, c) {
					continue
				}

				if a+b+c == sum {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
		c++
	}
	return a * b * c
}

// RunSolution9 finds the only Pythagorean triplet with a sum of 1000
func RunSolution9() int {
	return getProductOfPythagoreanTripletWithSum(1000)
}
