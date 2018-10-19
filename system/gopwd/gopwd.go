package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(args) == 1 || (len(args) == 2 && args[1] == "-L") {
		fmt.Println(pwd)
		os.Exit(0)
	}

	if len(args) == 2 && args[1] == "-P" {
		fi, err := os.Lstat(pwd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if fi.Mode()&os.ModeSymlink != 0 {
			realpath, err := filepath.EvalSymlinks(pwd)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(realpath)
			os.Exit(0)
		}

		os.Exit(0)
	}

	os.Exit(-1)
}

// -P, -PL, -LP pwd symlink
// -L pwd
// "" pwd
// -L a - too many arguments
// a -L - too many arguments
// -La - bad option
// -L -a - bad option
// -a -L - bad option
