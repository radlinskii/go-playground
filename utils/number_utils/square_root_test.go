package number_utils

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	var testsTable = []struct {
		name     string
		number   float64
		expected float64
	}{
		{"positive number #1", 13, math.Sqrt(13)},
		{"positive number #2", 31.1, math.Sqrt(31.1)},
		{"zero", 0, 0},
		{"negative number #1", -11.98, 0},
		{"negative number #2", -1, 0},
	}
	precision := 0.00000001

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got, _ := Sqrt(test.number)
			if test.number > 0 {
				if got < test.expected-precision || got > test.expected+precision {
					t.Errorf("\nSqrt(%g) \nexpected: %g \ngot: %g", test.number, test.expected, got)
				}
			} else {
				if got != test.expected {
					t.Errorf("\nSqrt(%g) \nexpected: %g \ngot: %g", test.number, test.expected, got)
				}
			}
		})
	}
}

func TestSqrt2(t *testing.T) {
	testsTable := []struct {
		name   string
		number float64
	}{
		{"positive number #1", 13},
		{"positive number #2", 31.1},
		{"zero", 0},
		{"negative number #1", -11.98},
		{"negative number #2", -1},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			_, err := Sqrt(test.number)

			if test.number >= 0 {
				if err != nil {
					t.Errorf("\nSqrt(%g): \nexpected: <nil> \ngot: %v", test.number, err)
				}
			} else {
				if err != ErrNegativeSqrt(test.number) {
					t.Errorf("\nSqrt(%g): \nexpected: %v \ngot: %v", test.number, ErrNegativeSqrt(test.number), err)
				}
			}
		})
	}
}

func TestErrNegativeSqrt_Error(t *testing.T) {
	testsTable := []struct {
		name     string
		number   float64
		expected string
	}{
		{"positive number #1", 13, "cannot Sqrt negative number: 13"},
		{"positive number #2", 31.1, "cannot Sqrt negative number: 31.1"},
		{"zero", 0, "cannot Sqrt negative number: 0"},
		{"negative number #1", -11.98, "cannot Sqrt negative number: -11.98"},
		{"negative number #2", -1, "cannot Sqrt negative number: -1"},
	}
	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := ErrNegativeSqrt(test.number).Error()
			if got != test.expected {
				t.Errorf("\nSqrt(%g): \nexpected: <nil> \ngot: %v", test.number, got)
			}
		})
	}
}
