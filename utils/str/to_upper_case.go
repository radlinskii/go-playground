package str

// ToUpperCase takes string and returns its copy transformed to contain only upper-letters.
func ToUpperCase(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i = i + 1 {
		if r[i] >= 97 && r[i] <= 122 {
			r[i] -= 32
		}
	}
	return string(r)
}
