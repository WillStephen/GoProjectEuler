package test

import (
	"GoProjectEuler/solutions"
	c "GoProjectEuler/test/common"
	"testing"
)

// TestMultiplesOf3And5 euofh
func TestMultiplesOf3And5(t *testing.T) {
	result := solutions.RunSolution1()

	c.AssertEqual(t, result, 233168, "")
}
