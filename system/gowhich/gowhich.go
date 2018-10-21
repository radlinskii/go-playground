package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	minusA := flag.Bool("a", false, "List all instances of executables found (instead of just the first one of each).")
	minusS := flag.Bool("s", false, "No output, just return 0 if all of the executables are found, or 1 if some were not found.")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("gowhich: too few arguments")
		os.Exit(1)
	}

	found := false
	path := os.Getenv("PATH")
	pathSlice := strings.Split(path, ":")
	for _, arg := range args {
		foundArg := false
		for _, dir := range pathSlice {
			fullPath := filepath.Join(dir, arg)
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				if mode.IsRegular() {
					if mode&0111 != 0 {
						found = true
						foundArg = true
						if !*minusS {
							fmt.Println(fullPath)
							if !*minusA {
								break
							}
						}
					}
				}
			}
		}
		if !foundArg {
			if !*minusA {
				found = false
			}
			if !*minusS && !*minusA {
				fmt.Printf("%s: not found\n", arg)
			}
		}
	}
	if found {
		os.Exit(0)
	}
	os.Exit(1)
}
