package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem14(t *testing.T) {
	result := solutions.RunSolution14()

	common.AssertEqual(t, 837799, result, "")
}

func BenchmarkProblem14(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution14()
	}
}
