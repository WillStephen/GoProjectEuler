package common

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func failTest(t *testing.T, message string) {
	t.Fatal(message)
	l := log.New(os.Stderr, "", 0)
	l.Println("Test failed.")
}

// AssertEqual asserts that two variables are equal
func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Expected: %v, Actual: %v", a, b)
	}
	failTest(t, message)
}
