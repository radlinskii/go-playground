package fibonacci

import "testing"

func TestMakeHeap(t *testing.T) {
	h := MakeHeap()

	if h.min != nil {
		t.Error("New FibonacciHeap should have nil as a pointer to the minimum element from its rootlist.")
	}

	if h.n != 0 {
		t.Error("New FibonacciHeap should have 0 as a number of elements.")
	}
}
