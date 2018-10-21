package algo

import (
	"reflect"
	"testing"
)

func TestSieve(t *testing.T) {
	var testsTable = []struct {
		name     string
		n        int
		expected []int
	}{
		{"zero", 0, []int{}},
		{"one", 1, []int{}},
		{"two", 2, []int{}},
		{"three", 3, []int{2}},
		{"four", 4, []int{2, 3}},
		{"five", 5, []int{2, 3}},
		{"seven", 7, []int{2, 3, 5}},
		{"eight", 8, []int{2, 3, 5, 7}},
		{"nine", 9, []int{2, 3, 5, 7}},
		{"nineteen", 19, []int{2, 3, 5, 7, 11, 13, 17}},
		{"twenty five", 25, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}},
		{"one hundred", 100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := Sieve(test.n)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("\nSieve(%d) \nexpected: %v \ngot: %v", test.n, test.expected, got)
			}
		})
	}
}
