package file

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestRemoveTabs(t *testing.T) {
	var testsTable = []struct {
		name, fileName string
	}{
		{"empty file", "remove_tabs1"},
		{"file without tabs", "remove_tabs2"},
		{"file with single line with a tab", "remove_tabs3"},
		{"file with multiple lines with tabs", "remove_tabs4"},
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

			// Reading input from .input file
			tmpInput, err := ioutil.ReadFile(input)
			if err != nil {
				t.Fatalf("failed reading input file: %s", err)
			}

			// Using RemoveTabs function to replace tab characters with 4 spaces
			err = RemoveTabs(input)
			if err != nil {
				t.Fatalf("failed removing tabs: %s", err)
			}

			// Reading output from .input file
			actual, err := ioutil.ReadFile(input)
			if err != nil {
				t.Fatalf("failed reading output: %s", err)
			}

			// Testing if expected output matches actual output
			if !bytes.Equal(actual, expected) {
				t.Errorf("\nCopy(%s) \nexpected: %s \ngot: %s", input, expected, actual)
			}

			// Writing expected output back to .output file
			err = ioutil.WriteFile(input, tmpInput, 0644)
			if err != nil {
				t.Fatalf("failed writing input back to input file: %s", err)
			}
		})
	}
}
