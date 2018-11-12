package algo

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	var testsTable = []struct {
		name     string
		curr     int
		new      int
		number   string
		expected string
	}{
		{"zero", 2, 8, "0", "0"},
		{"one", 2, 8, "1", "1"},
		{"555", 10, 8, "555", "1053"},
		{"30 binary to decimal", 2, 10, "11110", "30"},
		{"30 decimal to hexagonal", 10, 16, "30", "1E"},
		{"30 binary to hexagonal", 2, 16, "11110", "1E"},
		{"30 hexagonal to octal", 16, 8, "1E", "36"},
		{"30 hexagonal to decimal", 16, 10, "1E", "30"},
		{"30 decimal to octal", 10, 8, "30", "36"},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := Convert(test.curr, test.new, test.number)
			if got != test.expected {
				t.Errorf("\nexpected: %v \ngot: %v", test.expected, got)
			}
		})
	}
}

func TestToDecimal(t *testing.T) {
	var testsTable = []struct {
		name     string
		curr     int
		nums     []int
		expected int
	}{
		{"zero", 2, []int{0}, 0},
		{"one", 2, []int{1}, 1},
		{"binary two", 2, []int{1, 0}, 2},
		{"hexagonal 30", 16, []int{1, 14}, 30},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := toDecimal(test.curr, test.nums)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("\nexpected: %v \ngot: %v", test.expected, got)
			}
		})
	}
}

func TestFromDecimal(t *testing.T) {
	var testsTable = []struct {
		name     string
		new      int
		dec      int
		expected string
	}{
		{"zero", 2, 0, "0"},
		{"one", 2, 1, "1"},
		{"binary two", 2, 2, "01"},
		{"binary nine", 2, 9, "1001"},
		{"hexagonal thirty", 16, 30, "E1"},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := fromDecimal(test.new, test.dec)
			if got != test.expected {
				t.Errorf("\nexpected: %v \ngot: %v", test.expected, got)
			}
		})
	}
}
