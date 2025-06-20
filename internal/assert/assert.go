package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got %v; want %v", actual, expected)
	}
}

func StringContains(t *testing.T, actual, expectedSubstring string) {
	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf("got %v; want %v", actual, expectedSubstring)
	}
}
