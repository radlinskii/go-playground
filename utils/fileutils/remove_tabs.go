package fileutils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// RemoveTabs removes tabulation character from specified file and replaces them with 4 space character.
func RemoveTabs(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	output := strings.Replace(string(input), "\t", "    ", -1)

	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
	return
}
