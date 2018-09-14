package main

import (
	"os"

	"github.com/radlinskii/go-playground/utils/file_utils"
)

func main() {
	if len(os.Args) == 2 {
		file_utils.Write(os.Args[1])
	}

	file_utils.RemoveTabs("./file_utils/file.txt")
}
