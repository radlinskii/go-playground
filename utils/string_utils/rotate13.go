package string_utils

import "io"

// Rot13Reader is a simple structure that implements Reader interface.
// While read its encoding its content with ROT13 cipher.
type Rot13Reader struct {
	R io.Reader
}

func rotate13(b byte) byte {
	if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
		b += 13
	} else if (b > 'M' && b <= 'Z') || (b > 'm' && b <= 'z') {
		b -= 13
	}
	return b
}

// Read method is provided to implicitly implement the Reader interface to Rot13Reader type.
func (r13 *Rot13Reader) Read(s []byte) (n int, err error) {
	n, err = r13.R.Read(s)

	for i := 0; i < len(s); i++ {
		s[i] = rotate13(s[i])
	}
	return
}
