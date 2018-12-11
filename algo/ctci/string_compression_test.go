package ctci

import (
	"testing"
)

func TestCompress(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected string
	}{
		{"empty string", "", ""},
		{"string with 1 char", "a", "a"},
		{"string with 2 chars", "aa", "aa"},
		{"string with 2 chars", "aa", "aa"},
		{"string with 3 chars", "aaa", "a3"},
		{"not optimal string for compression", "aabbbca", "aabbbca"},
		{"string to be compressed", "aaaabbbbbca", "a4b5c1a1"},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := Compress(test.string)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
