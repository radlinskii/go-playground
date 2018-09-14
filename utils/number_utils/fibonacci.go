package number_utils

import "fmt"

func fibonacciWithClosure() func() int {
	i, a, b := 0, 0, 1
	return func() int {
		i = a
		a = a + b
		b = i
		return i
	}
}

// PrintFibonacciWithClosure prints first 10 Fibonacci numbers, implemented using closures.
func PrintFibonacciWithClosure() {
	f := fibonacciWithClosure()
	fmt.Print("first 10 of Fibonacci numbers: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
	fmt.Print("\n")
}

func fibonacciWithChannel(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// PrintFibonacciWithChannel prints first 10 Fibonacci numbers, implemented using goroutine & channel.
func PrintFibonacciWithChannel() {
	c := make(chan int, 10)
	go fibonacciWithChannel(cap(c), c)

	fmt.Print("first 10 of Fibonacci numbers: ")
	for i := range c {
		fmt.Printf("%d ", i)
	}
	fmt.Print("\n")
}
