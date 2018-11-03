package str

// Reverse takes a string and returns its reversed copy.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReverseRecursively takes a string and returns its reversed copy. Recursively.
func ReverseRecursively(s string) string {
	if len(s) <= 1 {
		return s
	}
	return ReverseRecursively(s[1:len(s)]) + string(s[0])
}
