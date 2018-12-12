package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

// TestProblem2
func TestProblem2(t *testing.T) {
	result := solutions.RunSolution2()

	common.AssertEqual(t, 4613732, result, "")
}
