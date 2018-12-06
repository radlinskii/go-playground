package ctci

// IsPalindromePermutation1 checks if string str is permuatation of a palindrome.
func IsPalindromePermutation1(str string) bool {
	var charMap [256]int
	for _, c := range str {
		charMap[int(c)]++
	}

	foundOneOdd := false
	for _, i := range charMap {
		if i%2 != 0 {
			if !foundOneOdd {
				foundOneOdd = true
			} else {
				return false
			}
		}
	}
	return true
}

// IsPalindromePermutation2 checks if string str is permuatation of a palindrome.
func IsPalindromePermutation2(str string) bool {
	var charMap [256]int
	countOdd := 0
	for _, c := range str {
		charMap[int(c)]++
		if charMap[int(c)]%2 != 0 {
			countOdd++
		} else {
			countOdd--
		}
	}
	return countOdd <= 1
}
