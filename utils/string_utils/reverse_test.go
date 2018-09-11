package string_utils

import "testing"

func TestReverse(t *testing.T) {
	var tests = map[string]string{
		"":     "",
		"aaa":  "aaa",
		"abc":  "cba",
		" ba ": " ab ",
	}

	for key, expectedVal := range tests {
		val := Reverse(key)
		if val != expectedVal {
			t.Error(
				"For", key,
				"expected", expectedVal,
				"got", val,
			)
		}
	}
}
