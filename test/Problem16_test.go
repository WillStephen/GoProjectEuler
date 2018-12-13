package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem16(t *testing.T) {
	result := solutions.RunSolution16()

	common.AssertEqual(t, 837799, result, "")
}

func BenchmarkProblem16(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution16()
	}
}
