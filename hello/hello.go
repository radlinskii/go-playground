package main

import (
	"fmt"
	"github.com/radlinskii/go-playground/string_utils"
	"github.com/radlinskii/go-playground/number_utils"
)

func main() {
	fmt.Printf("reversed message: %s\n", string_utils.Reverse("\n!oG ,olleH"))
	fmt.Printf("uppercase message: %s\n", string_utils.ToUpperCase("Hello, Go!\n"))

	var num float64 = 3
	sqrt, diff := number_utils.Sqrt(num)
	fmt.Printf("my sqrt of %g: %g\ndifference between my sqrt and math.Sqrt: %g\n", num, sqrt, diff)
}