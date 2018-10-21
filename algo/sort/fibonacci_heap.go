package sort

import (
	"github.com/radlinskii/go-playground/datastructures/heaps/fibonacci"
)

// FibonacciHeap is sorting a given slice using the Fibonacci Heap data structure.
func FibonacciHeap(s []int) []int {
	h := fibonacci.MakeHeap()
	for _, v := range s {
		h.Insert(fibonacci.MakeNode(v))
	}

	ns := []int{}
	for i := 0; i < len(s); i++ {
		node := h.ExtractMin()
		ns = append(ns, node.GetKey())
	}

	return ns
}
