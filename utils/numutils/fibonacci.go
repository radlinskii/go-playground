package numutils

import "fmt"

type errNegativeAmount float64

func (e errNegativeAmount) Error() string {
	return fmt.Sprintf("cannot operate on negative amount of numbers: %g", float64(e))
}

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
func FibonacciWithClosure(n int) ([]int, error) {
	if n < 0 {
		return []int{}, errNegativeAmount(n)
	}
	if n == 0 {
		return []int{}, nil
	}
	var s []int

	f := getFibonacciWithClosure()

	for i := 0; i < n; i++ {
		s = append(s, f())
	}
	return s, nil
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
func FibonacciWithChannel(n int) ([]int, error) {
	if n < 0 {
		return []int{}, errNegativeAmount(n)
	}
	if n == 0 {
		return []int{}, nil
	}
	var s []int
	c := make(chan int, n)

	go getFibonacciWithChannel(cap(c), c)

	for i := range c {
		s = append(s, i)
	}
	return s, nil
}
