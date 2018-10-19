package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// gopwd is an implementation of pwd(1) command.
func main() {
	minusL := flag.Bool("L", false, "displays the current logical working directory")
	minusP := flag.Bool("P", false, "displays the current physical working directory(all symbolic links resolved)")
	flag.Parse()

	if len(flag.Args()) != 0 {
		fmt.Println("gopwd: too many arguments")
		os.Exit(1)
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *minusL || !*minusP {
		fmt.Println(pwd)
		os.Exit(0)
	}

	if *minusP {
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

		fmt.Println(pwd)
		os.Exit(0)
	}

	os.Exit(1)
}
