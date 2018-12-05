package ctci

import (
	"testing"
)

func TestURLify(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected string
	}{
		{"empty string", "halo halo halo    ", "halo%20halo%20halo"},
		{"string of length 1", "   ", "%20"},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := URLify(test.string)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
