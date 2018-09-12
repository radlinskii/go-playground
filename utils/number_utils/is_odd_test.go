package number_utils

import "testing"

func TestIsOdd(t *testing.T) {
	var testsTable = []struct {
		name     string
		number   int
		expected bool
	}{
		{"positive odd number", 13, true},
		{"positive even number", 32, false},
		{"zero", 0, false},
		{"negative odd number", -11, true},
		{"negative even number", -44, false},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsOdd(test.number)
			if got != test.expected {
				t.Errorf("\nIsOdd(%d) \nexpected: %t \ngot: %t", test.number, got, test.expected)
			}

		})
	}
}
