package string_utils

import (
	"strings"
	"testing"
)

func ExamplePrintRot13Encoded() {
	PrintRot13Encoded("")
	PrintRot13Encoded("\n")
	PrintRot13Encoded("hello\n")
	PrintRot13Encoded("HELLO\n")
	PrintRot13Encoded("Hello!\n")
	// Output:
	// uryyb
	// URYYB
	// Uryyb!
}

func TestRot13Reader_Read(t *testing.T) {
	var testsTable = []struct {
		name     string
		r        *strings.Reader
		expected int
	}{
		{"empty strings.Reader", &strings.Reader{}, 0},
		{"strings.Reader with word", strings.NewReader("hello"), 5},
		{"strings.Reader with two words", strings.NewReader("hello world"), 11},
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
