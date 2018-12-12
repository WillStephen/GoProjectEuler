package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

// TestProblem1
func TestProblem1(t *testing.T) {
	result := solutions.RunSolution1()

	common.AssertEqual(t, 233168, result, "")
}
