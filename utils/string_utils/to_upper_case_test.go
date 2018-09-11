package string_utils

import "testing"

func TestToUpperCase(t *testing.T) {
	var tests = map[string]string{
		"":        "",
		"aaa":     "AAA",
		"ab!1$5c": "AB!1$5C",
		" ba ":    " BA ",
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
