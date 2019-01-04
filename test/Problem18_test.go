package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem18(t *testing.T) {
	result := solutions.RunSolution18()

	common.AssertEqual(t, 1074, result, "")
}

func BenchmarkProblem18(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution18()
	}
}
