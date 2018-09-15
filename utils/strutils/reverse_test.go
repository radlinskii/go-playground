package strutils

import "testing"

func TestReverse(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected string
	}{
		{"empty string", "", ""},
		{"word", "test", "tset"},
		{"palindrome", "level", "level"},
		{"string with spaces", "fat cat", "tac taf"},
		{"string with non-alphabetic characters", "#1337!!1", "1!!7331#"},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := Reverse(test.string)
			if got != test.expected {
				t.Errorf("\nReverse(%s) \nexpected: %s \ngot: %s", test.string, test.expected, got)
			}
		})
	}
}
