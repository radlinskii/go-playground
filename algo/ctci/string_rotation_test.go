package ctci

import "testing"

func TestIsRotation(t *testing.T) {
	var testsTable = []struct {
		name     string
		s1, s2   string
		expected bool
	}{
		{"empty string", "", "", false},
		{"string with 1 char", "a", "a", true},
		{"reversed string", "ab", "ba", true},
		{"not rotated string", "watermellon", "mellwateron", false},
		{"mellonwater", "watermellon", "mellonwater", true},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsRotation(test.s1, test.s2)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
