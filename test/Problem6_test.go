package test

import (
	"GoProjectEuler/common"
	"GoProjectEuler/solutions"
	"testing"
)

func TestProblem6(t *testing.T) {
	result := solutions.RunSolution6()

	common.AssertEqual(t, 25164150, result, "")
}

func BenchmarkProblem6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution6()
	}
}
