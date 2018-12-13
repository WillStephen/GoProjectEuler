package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

func TestProblem12(t *testing.T) {
	result := solutions.RunSolution12()

	common.AssertEqual(t, 70600674, result, "")
}

func BenchmarkProblem12(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution12()
	}
}
