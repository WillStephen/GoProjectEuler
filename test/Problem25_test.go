package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem25(t *testing.T) {
	result := solutions.RunSolution25()

	common.AssertEqual(t, 4782, result, "")
}

func BenchmarkProblem25(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution25()
	}
}
