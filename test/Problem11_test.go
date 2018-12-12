package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem11(t *testing.T) {
	result := solutions.RunSolution11()

	common.AssertEqual(t, 70600674, result, "")
}

func BenchmarkProblem11(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution11()
	}
}
