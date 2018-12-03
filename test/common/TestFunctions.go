package common

import (
	"fmt"
	"testing"
)

// AssertEqual asserts that two variables are equal
func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Expected: %v, Actual: %v", a, b)
	}
	t.Fatal(message)
}
