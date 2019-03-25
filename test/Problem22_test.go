package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem20(t *testing.T) {
	result := solutions.RunSolution22()

	common.AssertEqual(t, 871198282, result, "")
}

func BenchmarkProblem20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution22()
	}
}
