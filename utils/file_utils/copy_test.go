package file_utils

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
			golden := filepath.Join("testdata", test.fileName+".golden")

			// Reading expected output from .golden file
			expected, err := ioutil.ReadFile(golden)
			if err != nil {
				t.Fatalf("failed reading golden file: %s", err)
			}

			// Using Copy function to copy content of .input file to .golden file
			_, err = Copy(input, golden)
			if err != nil {
				t.Fatalf("failed copying file: %s", err)
			}

			// Reading output from .golden file
			actual, err := ioutil.ReadFile(golden)
			if err != nil {
				t.Fatalf("failed reading output: %s", err)
			}

			// Testing if expected output matches actual output
			if !bytes.Equal(actual, expected) {
				t.Errorf("\nCopy(%s) \nexpected: %s \ngot: %s", input, expected, actual)
			}

			// Writing expected output back to .golden file
			err = ioutil.WriteFile(golden, expected, 0644)
			if err != nil {
				t.Fatalf("failed writing expected output back to golden file: %s", err)
			}
		})
	}
}
