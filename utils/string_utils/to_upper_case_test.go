package string_utils

import "testing"

func TestToUpperCase(t *testing.T) {
	var tests = map[string]string{
		"":               "",
		"aaa":            "AAA",
		"string":         "STRING",
		" hello tests! ": " HELLO TESTS! ",
	}

	for key, expectedVal := range tests {
		val := ToUpperCase(key)
		if val != expectedVal {
			t.Error(
				"For", key,
				"expected", expectedVal,
				"got", val,
			)
		}
	}
}
