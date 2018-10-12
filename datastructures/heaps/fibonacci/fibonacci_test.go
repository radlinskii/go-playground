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

func TestMakeNode(t *testing.T) {
	expected := 4
	n := MakeNode(expected)

	if n.key != expected {
		t.Errorf("Expected Node Key: %d, got %d", expected, n.key)
	}

	if n.parent != nil || n.child != nil || n.left != nil || n.right != nil {
		t.Error("New Node should have nil as its pointers")
	}

	if n.mark != false {
		t.Error("New Node should be unmarked")
	}

	if n.degree != 0 {
		t.Errorf("New Node degree should be 0, got: %d", n.degree)
	}
}
