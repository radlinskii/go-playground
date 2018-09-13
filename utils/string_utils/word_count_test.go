package string_utils

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected map[string]int
	}{
		{"empty string", "", map[string]int{}},
		{"word", "test", map[string]int{"test": 1}},
		{"string with distinctive", "three different words", map[string]int{"three": 1, "different": 1, "words": 1}},
		{"string with repeating words", "cat dog mouse cat cat mouse", map[string]int{"cat": 3, "mouse": 2, "dog": 1}},
		{"string with same words in different cases", "cat Cat CAT", map[string]int{"cat": 1, "Cat": 1, "CAT": 1}},
		{"string with non-alphabetic characters", "Hello World! Hello Go!", map[string]int{"Hello": 2, "World!": 1, "Go!": 1}},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := WordCount(test.string)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("\nWordCount(%s) \nexpected: %v \ngot: %v", test.string, test.expected, got)
			}
		})
	}
}
