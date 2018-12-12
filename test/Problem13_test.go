package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem13(t *testing.T) {
	result := solutions.RunSolution13()

	common.AssertEqual(t, 5537376230, result, "")
}

func BenchmarkProblem13(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution13()
	}
}
