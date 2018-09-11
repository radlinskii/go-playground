package file_utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func RemoveTabs(filename string) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	output := strings.Replace(string(input), "\t", "    ", -1)

	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
