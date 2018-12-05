package test

import (
	"GoProjectEuler/common"
	"GoProjectEuler/solutions"
	"testing"
)

// TestProblem2
func TestProblem2(t *testing.T) {
	result := solutions.RunSolution2()

	common.AssertEqual(t, 4613732, result, "")
}
