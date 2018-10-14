package num

import "fmt"

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

// Sqrt calculates the square root of a number.
// If given negative number it returns an error.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errNegativeSqrt(x)
	} else if x == 0 {
		return 0, nil
	} else {
		z := float64(x)
		for i := 0; i < 100; i++ {
			z -= (z*z - x) / (2 * z)
		}
		return z, nil
	}
}
