package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem9(t *testing.T) {
	result := solutions.RunSolution9()

	common.AssertEqual(t, 31875000, result, "")
}

func BenchmarkProblem9(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution9()
	}
}
