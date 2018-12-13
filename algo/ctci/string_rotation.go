package ctci

import (
	"strings"
)

// IsRotation checks if string s2 is rotated string s1.
func IsRotation(s1, s2 string) bool {
	if len(s1) != len(s2) || len(s1) == 0 {
		return false
	}
	return strings.Contains(s1+s1, s2)
}
