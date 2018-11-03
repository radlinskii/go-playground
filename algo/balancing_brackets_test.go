package algo

import (
	"reflect"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected bool
	}{
		{"empty string", "", true},
		{"string of length 1", "[", false},
		{"string containing one bracket", "a}", false},
		{"square brackets", "[]", true},
		{"round brackets", "()", true},
		{"curly brackets", "{}", true},
		{"only opening brackets", "{([", false},
		{"only closing brackets", "])}", false},
		{"balanced brackets next to eachother", "{}()[]", true},
		{"balanced brakcets inside eachother", "{([])}", true},
		{"unbalanced brackets", "{[(])}", false},
		{"balanced combination", "{()[](([]))}", true},
		{"string containing a balanced brackets", "{a}([b]c)", true},
		{"string containing a unbalanced brackets", "{a)[[b]c}", false},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsBalanced(test.string)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("\nIsBalanced(%v) \nexpected: %v \ngot: %v", test.string, test.expected, got)
			}
		})
	}
}
