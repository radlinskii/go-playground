package main

import (
	"fmt"
	"flag"
	"github.com/radlinskii/go-playground/string_utils"
	"github.com/radlinskii/go-playground/number_utils"
	"github.com/radlinskii/go-playground/file_utils"
)

func main() {
	fmt.Printf("reversed message: %s\n", string_utils.Reverse("!oG ,olleH"))
	fmt.Printf("uppercase message: %s\n", string_utils.ToUpperCase("Hello, Go!"))

	var num float64 = 3
	sqrt, diff := number_utils.Sqrt(num)
	fmt.Printf("my sqrt of %g: %g\ndifference between my sqrt and math.Sqrt: %g\n", num, sqrt, diff)
	fmt.Printf("is %d a odd number? %t\n", int(num), number_utils.IsOdd(int(num)))
	file_utils.Copy("../file_utils/new_file.txt","../file_utils/file.txt")
	fmt.Println(string_utils.WordCount("Hello World!! World! I said Hello !!\n"))
	number_utils.PrintFibonacci()

	var name string
	flag.StringVar(&name, "name", "", "Who should i greet?")
    flag.Parse()

    if name == "" {
        fmt.Println("Hello World!")
    } else {
        fmt.Printf("Hello %s!\n", name)
    }
}
