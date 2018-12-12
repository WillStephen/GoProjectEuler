package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

// TestProblem5
func TestProblem5(t *testing.T) {
	result := solutions.RunSolution5()

	common.AssertEqual(t, 232792560, result, "")
}

func BenchmarkProblem5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution5()
	}
}
