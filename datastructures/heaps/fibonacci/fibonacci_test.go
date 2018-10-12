package fibonacci

import "testing"

func TestMakeHeap(t *testing.T) {
	h := MakeHeap()

	if h.min != nil {
		t.Error("New FibonacciHeap should have nil as a pointer to the minimum element from its rootlist")
	}

	if h.n != 0 {
		t.Error("New FibonacciHeap should have 0 as a number of elements")
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

func TestGetKey(t *testing.T) {
	expected := 4
	n := MakeNode(expected)
	actual := n.GetKey()

	if actual != expected {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}

	n = &Node{}
	actual = n.GetKey()
	if actual != 0 {
		t.Errorf("Empty node's key should be equal to 0, got: %d", actual)
	}
}

func TestAddNode(t *testing.T) {
	n1 := MakeNode(1)
	n1.left = n1
	n1.right = n1
	n2 := MakeNode(2)
	n2.left = n2
	n2.right = n2

	addNode(n1, n2)
	if n1.left != n2 || n2.right != n1 {
		t.Error("AddNode should add second node on the left of the first one")
	}
}
