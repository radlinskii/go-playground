package ctci

// IsOneEditAway checks if strings differ by zero or one editory operations.
func IsOneEditAway(str1, str2 string) bool {
	if len(str1) == len(str2) {
		foundOneEdit := false
		for i, c := range str1 {
			if str2[i] != byte(c) {
				if foundOneEdit {
					return false
				}
				foundOneEdit = true
			}
		}
		return true
	}
	if len(str1) == len(str2)+1 || len(str1)+1 == len(str2) {
		if len(str1) < len(str2) {
			str1, str2 = str2, str1
		}
		foundOneEdit := false
		shorterI, longerI := 0, 0
		for shorterI < len(str2) && longerI < len(str1) {
			if str1[longerI] == str2[shorterI] {
				shorterI++
			} else {
				if foundOneEdit {
					return false
				}
				foundOneEdit = true
			}
			longerI++
		}
		return true
	}
	return false
}
