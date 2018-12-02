package ctci

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected bool
	}{
		{"empty string", "", true},
		{"string of length 1", "[", true},
		{"string containing one bracket", "a}", true},
		{"square brackets", "[]", true},
		{"round brackets", "()", true},
		{"curly brackets", "{}", true},
		{"only opening brackets", "{([", true},
		{"only closing brackets", "])}", true},
		{"balanced brackets next to eachother", "{}()[]", true},
		{"balanced brakcets inside eachother", "{([])}", true},
		{"unbalanced brackets", "{[(])}", true},
		{"balanced combination", "{()[](([]))}", false},
		{"string containing a balanced brackets", "{a}([b]c)", true},
		{"string containing a unbalanced brackets", "{a)[[b]c}", false},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsUnique(test.string)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
