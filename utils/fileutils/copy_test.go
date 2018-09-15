package fileutils

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestCopy(t *testing.T) {
	var testsTable = []struct {
		name, fileName string
	}{
		{"empty file", "copy1"},
		{"file with single line of text", "copy2"},
		{"file with multiple lines of text", "copy3"},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			input := filepath.Join("testdata", test.fileName+".input")
			output := filepath.Join("testdata", test.fileName+".output")

			// Reading expected output from .output file
			expected, err := ioutil.ReadFile(output)
			if err != nil {
				t.Fatalf("failed reading output file: %s", err)
			}

			// Using Copy function to copy content of .input file to .output file
			_, err = Copy(input, output)
			if err != nil {
				t.Fatalf("failed copying file: %s", err)
			}

			// Reading output from .output file
			actual, err := ioutil.ReadFile(output)
			if err != nil {
				t.Fatalf("failed reading output: %s", err)
			}

			// Testing if expected output matches actual output
			if !bytes.Equal(actual, expected) {
				t.Errorf("\nCopy(%s) \nexpected: %s \ngot: %s", input, expected, actual)
			}

			// Writing expected output back to .output file
			err = ioutil.WriteFile(output, expected, 0644)
			if err != nil {
				t.Fatalf("failed writing expected output back to output file: %s", err)
			}
		})
	}
}
