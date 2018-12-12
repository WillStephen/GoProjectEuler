package test

import (
	"WillStephen/GoProjectEuler/common"
	"WillStephen/GoProjectEuler/solutions"
	"testing"
)

// TestProblem3
func TestProblem3(t *testing.T) {
	result := solutions.RunSolution3()

	common.AssertEqual(t, 6857, result, "")
}
