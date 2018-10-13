package sort

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	var testsTable = []struct {
		name     string
		slice    []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}},
		{"slice with single item", []int{3}, []int{3}},
		{"slice with sorted items", []int{-1, 32, 45}, []int{-1, 32, 45}},
		{"slice with unsorted items #1", []int{32, 0}, []int{0, 32}},
		{"slice with unsorted items #2", []int{94, 31, -3, 0}, []int{-3, 0, 31, 94}},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := Bubble(test.slice)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("\nBubble(%v) \nexpected: %v \ngot: %v", test.slice, test.expected, got)
			}
		})
	}
}
