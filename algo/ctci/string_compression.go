package ctci

import (
	"strconv"
)

// Compress compresses a string.
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
