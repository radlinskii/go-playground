package ctci

import (
	"reflect"
	"testing"
)

func TestLeftRotate(t *testing.T) {
	var testsTable = []struct {
		name     string
		arr      [][]int
		expected [][]int
	}{
		{"0x0", [][]int{}, [][]int{}},
		{"1x1", [][]int{{7}}, [][]int{{7}}},
		{"2x2", [][]int{{7, 3}, {5, 1}}, [][]int{{5, 7}, {1, 3}}},
		{"3x3", [][]int{{1, 2, 3}, {5, 6, 1}, {4, 1, 4}}, [][]int{{4, 5, 1}, {1, 6, 2}, {4, 1, 3}}},
		{"4x4", [][]int{{1, 2, 3, 6}, {5, 6, 1, 5}, {4, 1, 4, 6}, {4, 7, 1, 8}}, [][]int{{4, 4, 5, 1}, {7, 1, 6, 2}, {1, 4, 1, 3}, {8, 6, 5, 6}}},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			LeftRotate(test.arr)
			if !reflect.DeepEqual(test.arr, test.expected) {
				t.Errorf("expected: %v, got: %v", test.expected, test.arr)
			}
		})
	}
}
