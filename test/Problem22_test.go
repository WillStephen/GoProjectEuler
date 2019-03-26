package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem22(t *testing.T) {
	result := solutions.RunSolution22()

	common.AssertEqual(t, 871198282, result, "")
}

func BenchmarkProblem22(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution22()
	}
}
