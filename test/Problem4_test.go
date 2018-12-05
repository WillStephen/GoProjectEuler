package test

import (
	"GoProjectEuler/common"
	"GoProjectEuler/solutions"
	"testing"
)

// TestProblem4
func TestProblem4(t *testing.T) {
	result := solutions.RunSolution4()

	common.AssertEqual(t, 906609, result, "")
}
