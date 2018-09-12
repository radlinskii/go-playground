package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/radlinskii/go-playground/utils/file_utils"
	"github.com/radlinskii/go-playground/utils/string_utils"
)

func main() {
	file_utils.Copy("./file_utils/new_file.txt", "./file_utils/file.txt")

	if len(os.Args) == 2 {
		file_utils.Write(os.Args[1])
	}

	ipAddr := string_utils.IPAddr{192, 168, 182, 157}
	fmt.Printf("stringified ip address type: %v\n", ipAddr)

	s := strings.NewReader("Uryyb Jbeyq!\n")
	r := string_utils.Rot13Reader{s}
	fmt.Print("decoded message: ")
	io.Copy(os.Stdout, &r)

	file_utils.RemoveTabs("./file_utils/file.txt")
}
