package ctci

import (
	"reflect"
	"testing"
)

func TestIsPermutation(t *testing.T) {
	var testsTable = []struct {
		name     string
		str1     string
		str2     string
		expected bool
	}{
		{"empty string", "", "", true},
		{"string of length 1", "[", "[", true},
		{"string containing one bracket", "a}", "}a", true},
		{"string containing a balanced brackets", "{a}([b]c)", "{}abc()[]", true},
		{"string containing a unbalanced brackets", "{a)[[b]c}", "{}abc()[", false},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsPermutation(test.str1, test.str2)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}

func TestIsUnique(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected bool
	}{
		{"empty string", "", true},
		{"string of length 1", "[", true},
		{"string containing one bracket", "a}", true},
		{"square brackets", "[]", true},
		{"round brackets", "()", true},
		{"curly brackets", "{}", true},
		{"only opening brackets", "{([", true},
		{"only closing brackets", "])}", true},
		{"balanced brackets next to eachother", "{}()[]", true},
		{"balanced brakcets inside eachother", "{([])}", true},
		{"unbalanced brackets", "{[(])}", true},
		{"balanced combination", "{()[](([]))}", false},
		{"string containing a balanced brackets", "{a}([b]c)", true},
		{"string containing a unbalanced brackets", "{a)[[b]c}", false},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsUnique(test.string)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}

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

func TestIsOneEditAway(t *testing.T) {
	var testsTable = []struct {
		name       string
		str1, str2 string
		expected   bool
	}{
		{"empty string", "", "", true},
		//{"string of length 1", "[", "", true},
		//{"string containing one bracket", "a}", "a}a", true},
		{"square brackets", "[]", "53", false},
		{"replacement", "pale", "bale", true},
		{"double replacement", "pale", "bake", false},
		{"missing letter", "pale", "ale", true},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsOneEditAway(test.str1, test.str2)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}

func TestIsPalindromePermutation1(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected bool
	}{
		{"empty string", "", true},
		{"string of length 1", "[", true},
		{"string containing one bracket", "a}", false},
		{"balanced brackets next to eachother", "{}({}", true},
		{"balanced brakcets inside eachother", "{([{[(", true},
		{"string containing a balanced brackets", "{a}([b]c)", false},
		{"string containing a unbalanced brackets", "{a)[[b]c}", false},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsPalindromePermutation1(test.string)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}

func TestIsPalindromePermutation2(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected bool
	}{
		{"empty string", "", true},
		{"string of length 1", "[", true},
		{"string containing one bracket", "a}", false},
		{"balanced brackets next to eachother", "{}({}", true},
		{"balanced brakcets inside eachother", "{([{[(", true},
		{"string containing a balanced brackets", "{a}([b]c)", false},
		{"string containing a unbalanced brackets", "{a)[[b]c}", false},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := IsPalindromePermutation2(test.string)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}

func TestCompress(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected string
	}{
		{"empty string", "", ""},
		{"string with 1 char", "a", "a"},
		{"string with 2 chars", "aa", "aa"},
		{"string with 2 chars", "aa", "aa"},
		{"string with 3 chars", "aaa", "a3"},
		{"not optimal string for compression", "aabbbca", "aabbbca"},
		{"string to be compressed", "aaaabbbbbca", "a4b5c1a1"},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := Compress(test.string)
			if got != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
