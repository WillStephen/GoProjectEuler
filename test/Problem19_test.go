package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem19(t *testing.T) {
	result := solutions.RunSolution19()

	common.AssertEqual(t, 171, result, "")
}

func BenchmarkProblem19(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution19()
	}
}
