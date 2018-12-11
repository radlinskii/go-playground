package ctci

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	var testsTable = []struct {
		name     string
		arr      [][]int
		expected [][]int
	}{
		{"0x0", [][]int{}, [][]int{}},
		{"1x1", [][]int{{7}}, [][]int{{7}}},
		{"2x2", [][]int{{7, 0}, {5, 1}}, [][]int{{0, 0}, {5, 0}}},
		{"3x3", [][]int{{0, 2, 3}, {5, 0, 1}, {4, 1, 0}}, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		{"4x4", [][]int{{1, 2, 3, 0}, {5, 6, 1, 5}, {0, 1, 4, 6}, {4, 7, 1, 8}}, [][]int{{0, 0, 0, 0}, {0, 6, 1, 0}, {0, 0, 0, 0}, {0, 7, 1, 0}}},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			ZeroMatrix(test.arr)
			if !reflect.DeepEqual(test.arr, test.expected) {
				t.Errorf("expected: %v, got: %v", test.expected, test.arr)
			}
		})
	}
}
