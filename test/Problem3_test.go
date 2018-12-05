package test

import (
	"GoProjectEuler/common"
	"GoProjectEuler/solutions"
	"testing"
)

// TestProblem3
func TestProblem3(t *testing.T) {
	result := solutions.RunSolution3()

	common.AssertEqual(t, 6857, result, "")
}
