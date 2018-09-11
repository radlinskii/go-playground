package string_utils

import "testing"

func TestWordCount(t *testing.T) {
	var tests = map[string]map[string]int{
		"":                                    {},
		"word":                                {"word": 1},
		"a simple word ":                      {"a": 1, "simple": 1, "word": 1},
		" first word second word other word!": {"first": 1, "word": 2, "second": 1, "word!": 1, "other": 1},
	}

	for key, expectedMap := range tests {
		evaluatedMap := WordCount(key)
		for k, evaluatedVal := range evaluatedMap {
			if evaluatedVal != expectedMap[k] {
				t.Error(
					"For", key,
					"expected", expectedMap[k],
					"got", evaluatedVal,
				)
			}
		}

	}
}
