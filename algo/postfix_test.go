package algo

import (
	"testing"
)

func TestReversePolishNotation(t *testing.T) {
	var testsTable = []struct {
		name     string
		expr     string
		expected float32
	}{
		{"expressions with numbers lower than 10 #1", "2 3 + 5 *", 25},
		{"expressions with numbers lower than 10 #2", "2 3 1 * + 9 -", -4},
		{"expressions with numbers greater than 10 #1", "142 43 -", 99},
		{"expressions with numbers greater than 10 #2", "2 7 + 3 / 14 3 - 4 * + 2 /", 23.5},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := ReversePolishNotation(test.expr)
			if got != test.expected {
				t.Errorf("\nexpected: %v \ngot: %v", test.expected, got)
			}
		})
	}
}
