package number_utils

import "math"

func Sqrt(x float64) (float64, float64) {
	z := float64(x)
	for i := 0; i < 100; i++  {
		z -= (z*z - x) / (2 * z)
	}
	return z, math.Abs(z - math.Sqrt(x))
}
