package number_utils

import "testing"

func TestIsOdd(t *testing.T) {
	var tests = map[int]bool{
		3: true,
		2: false,
		1: true,
		0: false,
		4: false,
		5: true,
	}

	for key, expectedVal := range tests {
		val := IsOdd(key)
		if val != expectedVal {
			t.Error(
				"For", key,
				"expected", expectedVal,
				"got", val,
			)
		}
	}
}
