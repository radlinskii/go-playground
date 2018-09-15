package strutils

import (
	"strings"
	"testing"
)

func TestEncodeROT13(t *testing.T) {
	var testsTable = []struct {
		name string
		string
		expected string
	}{
		{"empty string", "", ""},
		{"word", "hello", "uryyb"},
		{"string with capital letters", "HellO", "UryyB"},
		{"string with non-alphabetic characters", "Hello World!", "Uryyb Jbeyq!"},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := EncodeROT13(test.string)
			if got != test.expected {
				t.Errorf("\nEncodeROT13(%v) \nexpected: %s \ngot: %s", test.string, test.expected, got)
			}
		})
	}
}

func TestRot13Reader_Read(t *testing.T) {
	var testsTable = []struct {
		name     string
		r        *strings.Reader
		expected int
	}{
		{"empty strings.Reader", &strings.Reader{}, 0},
		{"strings.Reader with word", strings.NewReader("hello"), 5},
		{"strings.Reader with capital letters", strings.NewReader("HellO"), 5},
		{"strings.Reader with non-alphabetic characters", strings.NewReader("Hello World!"), 12},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			s := make([]byte, 100)
			got, _ := (&rot13Reader{test.r}).Read(s)
			if got != test.expected {
				t.Errorf("\nRot13Reader_Read(%v) \nexpected: %d \ngot: %d", test.r, test.expected, got)
			}
		})
	}
}
