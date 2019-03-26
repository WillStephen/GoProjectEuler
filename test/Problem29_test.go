package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem29(t *testing.T) {
	result := solutions.RunSolution29()

	common.AssertEqual(t, 9183, result, "")
}

func BenchmarkProblem29(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution29()
	}
}
