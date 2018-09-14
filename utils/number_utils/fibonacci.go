package number_utils

// TODO handle n == 0 && n < 0

func getFibonacciWithClosure() func() int {
	i, a, b := 0, 0, 1
	return func() int {
		i = a
		a = a + b
		b = i
		return i
	}
}

// FibonacciWithClosure returns a slice of first n Fibonacci numbers
// implemented using closures.
func FibonacciWithClosure(n int) []int {
	s := make([]int, n)

	f := getFibonacciWithClosure()

	for i := 0; i < n; i++ {
		s = append(s, f())
	}
	return s
}

func getFibonacciWithChannel(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// FibonacciWithChannel returns a slice of first n Fibonacci numbers
// implemented using goroutine & channel.
func FibonacciWithChannel(n int) []int {
	s := make([]int, n)
	c := make(chan int, n)

	go getFibonacciWithChannel(cap(c), c)

	for i := range c {
		s = append(s, i)
	}
	return s
}
