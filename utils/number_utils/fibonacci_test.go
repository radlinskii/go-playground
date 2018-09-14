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
		{"negative amount of numbers #1", -24, []int{}},
		{"negative amount of numbers #2", -1, []int{}},
		{"no items", 0, []int{}},
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
			got, err := FibonacciWithClosure(test.amountOfNumbers)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("\nFibonacciWithClosure(%d) \nexpected: %v \ngot: %v", test.amountOfNumbers, test.expected, got)
			}
			if test.amountOfNumbers >= 0 {
				if err != nil {
					t.Errorf("\nFibonacciWithClosure(%d) \nexpected: %v \ngot: %v", test.amountOfNumbers, nil, err)
				}
			} else {
				if _, ok := err.(errNegativeAmount); !ok {
					t.Errorf("\nFibonacciWithClosure(%d) \nexpected: %v \ngot: %v",
						test.amountOfNumbers, errNegativeAmount(test.amountOfNumbers), err)
				}
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
		{"negative amount of numbers #1", -24, []int{}},
		{"negative amount of numbers #2", -1, []int{}},
		{"no items", 0, []int{}},
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
			got, err := FibonacciWithChannel(test.amountOfNumbers)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("\nFibonacciWithChannel(%d) \nexpected: %v \ngot: %v", test.amountOfNumbers, test.expected, got)
			}
			if test.amountOfNumbers >= 0 {
				if err != nil {
					t.Errorf("\nFibonacciWithChannel(%d) \nexpected: %v \ngot: %v", test.amountOfNumbers, nil, err)
				}
			} else {
				if _, ok := err.(errNegativeAmount); !ok {
					t.Errorf("\nFibonacciWithChannel(%d) \nexpected: %v \ngot: %v",
						test.amountOfNumbers, errNegativeAmount(test.amountOfNumbers), err)
				}
			}
		})
	}
}

func TestErrNegativeAmount_Error(t *testing.T) {
	testsTable := []struct {
		name     string
		number   float64
		expected string
	}{
		{"negative amount of numbers #1", -39, "cannot operate on negative amount of numbers: -39"},
		{"negative amount of numbers #2", -24, "cannot operate on negative amount of numbers: -24"},
		{"negative amount of numbers #3", -1, "cannot operate on negative amount of numbers: -1"},
		{"no items", 0, "cannot operate on negative amount of numbers: 0"},
		{"1 item", 1, "cannot operate on negative amount of numbers: 1"},
		{"7 items", 7, "cannot operate on negative amount of numbers: 7"},
	}
	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := errNegativeAmount(test.number).Error()
			if got != test.expected {
				t.Errorf("\nErrNegativeSqrt(%g).Error(): \nexpected: <nil> \ngot: %v", test.number, got)
			}
		})
	}
}
