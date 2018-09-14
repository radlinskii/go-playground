package number_utils

import (
	"reflect"
	"testing"
)

func TestFibonacciWithClosure(t *testing.T) {
	var testsTable = []struct {
		name            string
		amountOfNumbers int
		expected        []int
	}{
		{"no items", 0, nil},
		{"1 item", 1, []int{0}},
		{"2 items", 2, []int{0, 1}},
		{"3 items", 3, []int{0, 1, 1}},
		{"4 items", 4, []int{0, 1, 1, 2}},
		{"5 items", 5, []int{0, 1, 1, 2, 3}},
		{"6 items", 6, []int{0, 1, 1, 2, 3, 5}},
		{"7 items", 7, []int{0, 1, 1, 2, 3, 5, 8}},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := FibonacciWithClosure(test.amountOfNumbers)
			if reflect.DeepEqual(got, test.expected) {
				t.Errorf("\nFibonacciWithClosure(%d) \nexpected: %v \ngot: %v", test.amountOfNumbers, test.expected, got)
			}
		})
	}
}

func TestFibonacciWithChannel(t *testing.T) {
	var testsTable = []struct {
		name            string
		amountOfNumbers int
		expected        []int
	}{
		{"no items", 0, nil},
		{"1 item", 1, []int{0}},
		{"2 items", 2, []int{0, 1}},
		{"3 items", 3, []int{0, 1, 1}},
		{"4 items", 4, []int{0, 1, 1, 2}},
		{"5 items", 5, []int{0, 1, 1, 2, 3}},
		{"6 items", 6, []int{0, 1, 1, 2, 3, 5}},
		{"7 items", 7, []int{0, 1, 1, 2, 3, 5, 8}},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := FibonacciWithChannel(test.amountOfNumbers)
			if reflect.DeepEqual(got, test.expected) {
				t.Errorf("\nFibonacciWithChannel(%d) \nexpected: %v \ngot: %v", test.amountOfNumbers, test.expected, got)
			}
		})
	}
}
