package file_utils

import (
	"fmt"
	"os"
)

// Write writes "Hello Writing to Files!\n" string to a given file.
func Write(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b := []byte("Hello Writing to Files!\n")
	n, err := file.Write(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("number of bytes written: %d\n", n)
}
