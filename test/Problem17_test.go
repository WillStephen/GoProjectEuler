package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem17(t *testing.T) {
	result := solutions.RunSolution17()

	common.AssertEqual(t, 21124, result, "")
}

func BenchmarkProblem17(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution17()
	}
}
