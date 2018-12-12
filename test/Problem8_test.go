package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem8(t *testing.T) {
	result := solutions.RunSolution8()

	common.AssertEqual(t, 23514624000, result, "")
}

func BenchmarkProblem8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution8()
	}
}
