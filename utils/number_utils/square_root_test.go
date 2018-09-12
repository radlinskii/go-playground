package number_utils

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	var tests = map[float64]float64{
		3:  math.Sqrt(3),
		2:  math.Sqrt(2),
		1:  math.Sqrt(1),
		0:  math.Sqrt(0),
		4:  math.Sqrt(4),
		5:  math.Sqrt(5),
		-5: 0,
		-1: 0,
	}
	precision := 0.00000001

	for key, expectedVal := range tests {
		val, _ := Sqrt(key)
		if val < expectedVal-precision || val > expectedVal+precision {
			t.Error(
				"For", key,
				"expected", expectedVal,
				"got", val,
			)
		}
	}
}

func TestErrNegativeSqrt_Error(t *testing.T) {
	tests := [8]float64{-3.5, -3.4, -2, -1, 0, 1, 2, 5}

	for _, testNumber := range tests {
		_, err := Sqrt(testNumber)

		if testNumber < 0 {
			_, ok := err.(ErrNegativeSqrt)
			if !ok {
				t.Error(
					"For", testNumber,
					"expected ErrNegativeSqrt",
					"got", err,
				)
			}
		} else {
			if err != nil {
				t.Error(
					"For", testNumber,
					"expected <nil>",
					"got", err,
				)
			}
		}
	}
}
