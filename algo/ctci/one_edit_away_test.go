package ctci

import (
	"testing"
)

func TestIsOneEditAway(t *testing.T) {
	var testsTable = []struct {
		name       string
		str1, str2 string
		expected   bool
	}{
		{"empty string", "", "", true},
		//{"string of length 1", "[", "", true},
		//{"string containing one bracket", "a}", "a}a", true},
		{"square brackets", "[]", "53", false},
		{"replacement", "pale", "bale", true},
		{"double replacement", "pale", "bake", false},
		{"missing letter", "pale", "ale", true},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsOneEditAway(test.str1, test.str2)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
