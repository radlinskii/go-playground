package ctci

// URLify transforms a string to URL format. It returns new string.
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
