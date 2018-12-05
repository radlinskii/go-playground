package ctci

import (
	"testing"
)

func TestIsPermutation(t *testing.T) {
	var testsTable = []struct {
		name     string
		str1     string
		str2     string
		expected bool
	}{
		{"empty string", "", "", true},
		{"string of length 1", "[", "[", true},
		{"string containing one bracket", "a}", "}a", true},
		{"string containing a balanced brackets", "{a}([b]c)", "{}abc()[]", true},
		{"string containing a unbalanced brackets", "{a)[[b]c}", "{}abc()[", false},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsPermutation(test.str1, test.str2)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
