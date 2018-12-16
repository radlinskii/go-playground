package ctci

import (
	"strconv"
	"strings"
)

// IsUnique checks if all characters in a string are distinct.
// Exercise 1.1
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

// IsPermutation checks if string str1 is permutation of string str2.
// Exercise 1.2
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

// URLify transforms a string to URL format. It returns new string.
// Exercise 1.3
func URLify(str string) string {
	runes := []rune(str)
	for i, j := 0, 0; j < len(str); i++ {
		if j > len(str) {
			break
		}
		if str[i] != ' ' {
			runes[j] = rune(str[i])
			j++
		} else {
			runes[j] = '%'
			runes[j+1] = '2'
			runes[j+2] = '0'
			j = j + 3
		}
	}
	return string(runes)
}

// IsPalindromePermutation1 checks if string str is permuatation of a palindrome.
// Exercise 1.4
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
// Exercise 1.4
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

// IsOneEditAway checks if strings differ by zero or one editory operations.
// Exercise 1.5
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

// Compress compresses a string.
// Exercise 1.6
func Compress(str string) string {
	if len(str) < 3 {
		return str
	}
	c := str[0]
	count := 1
	res := string(c)
	for i := 1; i < len(str); i++ {
		if str[i] == c {
			count++
		} else {
			c = str[i]
			res += strconv.Itoa(count) + string(c)
			count = 1
		}
	}
	res += strconv.Itoa(count)
	if len(str) <= len(res) {
		return str
	}
	return res
}

// LeftRotate rotates the matrix by 90 degrees.
// Exercise 1.7
func LeftRotate(arr [][]int) {
	if len(arr) > 1 {
		for i := 0; i <= len(arr)/2; i++ {
			for j := i; j+i < len(arr)-1; j++ {
				tmp := arr[i][j]
				arr[i][j] = arr[len(arr)-1-j][i]
				arr[len(arr)-1-j][i] = arr[len(arr)-1-i][len(arr)-1-j]
				arr[len(arr)-1-i][len(arr)-1-j] = arr[j][len(arr)-1-i]
				arr[j][len(arr)-1-i] = tmp
			}
		}
	}
}

// ZeroMatrix modifies a matrix so that if it finds a zero in the matrix,
// it turns its whole row and column into zeros. Works for Matrixes of dimension MxN.
// Exercise 1.8
func ZeroMatrix(arr [][]int) {
	m := len(arr)
	if m > 0 {
		n := len(arr[0])

		rows := make([]bool, m)
		cols := make([]bool, n)

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if arr[i][j] == 0 {
					rows[i] = true
					cols[j] = true
				}
			}
		}
		// "zero out" rows
		for i := 0; i < m; i++ {
			if rows[i] == true {
				for j := 0; j < n; j++ {
					arr[i][j] = 0
				}
			}
		}
		// "zero out" columns
		for i := 0; i < n; i++ {
			if cols[i] == true {
				for j := 0; j < m; j++ {
					arr[j][i] = 0
				}
			}
		}
	}
}

// IsRotation checks if string s2 is rotated string s1.
// Exercise 1.9
func IsRotation(s1, s2 string) bool {
	if len(s1) != len(s2) || len(s1) == 0 {
		return false
	}
	return strings.Contains(s1+s1, s2)
}
