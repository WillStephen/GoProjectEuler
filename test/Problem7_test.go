package test

import (
	"GoProjectEuler/common"
	"GoProjectEuler/solutions"
	"testing"
)

func TestProblem7(t *testing.T) {
	result := solutions.RunSolution7()

	common.AssertEqual(t, 104743, result, "")
}

func BenchmarkProblem7(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution7()
	}
}
