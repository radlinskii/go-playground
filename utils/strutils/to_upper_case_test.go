package strutils

import "testing"

func TestToUpperCase(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected string
	}{
		{"empty string", "", ""},
		{"word", "test", "TEST"},
		{"uppercase word", "LEVEL", "LEVEL"},
		{"string with spaces", "fat cat", "FAT CAT"},
		{"string with non-alphabetic characters", "#leet!!1", "#LEET!!1"},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := ToUpperCase(test.string)
			if got != test.expected {
				t.Errorf("\nToUpperCase(%s) \nexpected: %s \ngot: %s", test.string, test.expected, got)
			}
		})
	}
}
