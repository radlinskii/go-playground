package ctci

// IsPermutation checks if string str1 is permutation of string str2.
func IsPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	var charMap [256]int
	for _, c := range str1 {
		charMap[int(c)]++
	}
	for _, c := range str2 {
		if charMap[int(c)] == 0 {
			return false
		}
		charMap[int(c)]--
	}
	for _, i := range charMap {
		if i != 0 {
			return false
		}
	}
	return true
}
