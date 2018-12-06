package ctci

import (
	"testing"
)

func TestIsPalindromePermutation1(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected bool
	}{
		{"empty string", "", true},
		{"string of length 1", "[", true},
		{"string containing one bracket", "a}", false},
		{"balanced brackets next to eachother", "{}({}", true},
		{"balanced brakcets inside eachother", "{([{[(", true},
		{"string containing a balanced brackets", "{a}([b]c)", false},
		{"string containing a unbalanced brackets", "{a)[[b]c}", false},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsPalindromePermutation1(test.string)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}

func TestIsPalindromePermutation2(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected bool
	}{
		{"empty string", "", true},
		{"string of length 1", "[", true},
		{"string containing one bracket", "a}", false},
		{"balanced brackets next to eachother", "{}({}", true},
		{"balanced brakcets inside eachother", "{([{[(", true},
		{"string containing a balanced brackets", "{a}([b]c)", false},
		{"string containing a unbalanced brackets", "{a)[[b]c}", false},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsPalindromePermutation2(test.string)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
