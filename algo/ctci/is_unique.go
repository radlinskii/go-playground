package ctci

// IsUnique checks if all characters in a string are distinct.
func IsUnique(str string) bool {
	if len(str) > 256 {
		return false
	}
	var charMap [256]bool
	for _, c := range str {
		if charMap[int(c)] == true {
			return false
		}
		charMap[int(c)] = true
	}
	return true
}
