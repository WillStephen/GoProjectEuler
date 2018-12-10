package test

import (
	"GoProjectEuler/common"
	"GoProjectEuler/solutions"
	"testing"
)

func TestProblem10(t *testing.T) {
	result := solutions.RunSolution10()

	common.AssertEqual(t, 142913828922, result, "")
}

func BenchmarkProblem10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solutions.RunSolution10()
	}
}
