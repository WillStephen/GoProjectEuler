package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem28(t *testing.T) {
	result := solutions.RunSolution28()

	common.AssertEqual(t, 669171001, result, "")
}

func BenchmarkProblem28(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution28()
	}
}
