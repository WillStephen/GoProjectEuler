package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem30(t *testing.T) {
	result := solutions.RunSolution30()

	common.AssertEqual(t, 443839, result, "")
}

func BenchmarkProblem30(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution30()
	}
}
