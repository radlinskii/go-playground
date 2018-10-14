package str

import (
	"bytes"
	"io"
	"strings"
)

type rot13Reader struct {
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
func (r13 rot13Reader) Read(s []byte) (n int, err error) {
	n, err = r13.R.Read(s)

	for i := 0; i < len(s); i++ {
		s[i] = rotate13(s[i])
	}
	return
}

// EncodeROT13 gets a string and returns its content encoded with ROT13 cipher
func EncodeROT13(s string) string {
	r := strings.NewReader(s)
	r13 := rot13Reader{r}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r13)
	return buf.String()
}
