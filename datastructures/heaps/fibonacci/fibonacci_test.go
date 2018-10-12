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

func TestAddRoot(t *testing.T) {
	h := MakeHeap()
	n1 := MakeNode(2)
	n2 := MakeNode(4)
	n3 := MakeNode(1)

	t.Run("addRoot on empty Fibonacci Heap", func(t *testing.T) {
		h.addRoot(n1)
		if h.min != n1 {
			t.Error("addRoot should make new node assigned to h.min")
		}
	})

	t.Run("addRoot with node's value bigger than h.min's key", func(t *testing.T) {
		h.addRoot(n2)
		if h.min == n2 {
			t.Error("addRoot should leave the old h.min")
		}

		if h.min.left != n2 {
			t.Error("addRoot should add the root on the left h.min")
		}
	})

	t.Run("addRoot with node's value lesser than the h.min's key", func(t *testing.T) {
		old := h.min
		h.addRoot(n3)

		if old.left != n3 {
			t.Error("addRoot should add the root on the left of previous h.min")
		}

		if h.min != n3 {
			t.Error("addRoot should make new node's assigned to h.min")
		}
	})
}

func TestRemoveNodeFromList(t *testing.T) {
	h := MakeHeap()
	n1 := MakeNode(2)
	n2 := MakeNode(4)

	h.addRoot(n1)
	h.addRoot(n2)

	removeNodeFromList(n1)

	if n2.left == n1 || n1.right == n1 {
		t.Error("removeNodeFromList should remove connections to node from the list it was in")
	}

	if n1.left != n2 || n1.right != n2 {
		t.Error("removeNodeFromList shouldn't modify node's pointers")
	}

	removeNodeFromList(n2)

	if n2.left != n2 || n2.right != n2 {
		t.Error("removeNodeFromList should't be able to remove node from list if it's the only element in the list")
	}
}

func TestInsert(t *testing.T) {
	h := MakeHeap()
	n := MakeNode(3)

	t.Run("Insert on empty Heap", func(t *testing.T) {
		n.degree = 2
		n.mark = true
		n.parent = n
		n.child = n

		h.Insert(n)

		if n.degree != 0 {
			t.Errorf("Node's degree should equal to 0, got %d", n.degree)
		}

		if n.mark != false {
			t.Error("Node should be unmarked")
		}

		if n.parent != nil {
			t.Error("Node's parent should be nil")
		}

		if n.child != nil {
			t.Error("Node's child should be nil")
		}

		if h.min != n {
			t.Error("Insert should make node assigned to h.min")
		}

		if h.n != 1 {
			t.Errorf("Heaps number of nodes should increment to 1, got %d", h.n)
		}
	})
}

func TestMinimum(t *testing.T) {
	h := MakeHeap()
	n1 := MakeNode(4)
	n2 := MakeNode(2)

	if h.Minimum() != nil {
		t.Error("Minimum on empty heap should return nil")
	}

	h.Insert(n1)

	if h.Minimum() != n1 {
		t.Error("Minimum should return h.min from heap with single element")
	}

	h.Insert(n2)

	if h.Minimum() != n2 {
		t.Error("Minimum should return h.min from updated heap")
	}
}

func TestUnion1(t *testing.T) {
	h1 := MakeHeap()
	n1 := MakeNode(4)
	n2 := MakeNode(2)
	h1.Insert(n1)
	h1.Insert(n2)

	h2 := MakeHeap()
	n3 := MakeNode(5)
	n4 := MakeNode(8)
	h2.Insert(n3)
	h2.Insert(n4)

	h3 := h1.Union(h2)

	if h3.n != h1.n+h2.n {
		t.Errorf("Merged heap's number of nodes should be %d, got %d", h1.n+h2.n, h3.n)
	}

	if h3.min != n2 {
		t.Error("Merged heap's min should be the minimum of two heaps")
	}
}

func TestUnion2(t *testing.T) {
	h1 := MakeHeap()
	n1 := MakeNode(4)
	n2 := MakeNode(2)
	h1.Insert(n1)
	h1.Insert(n2)

	h2 := MakeHeap()
	n3 := MakeNode(1)
	n4 := MakeNode(8)
	h2.Insert(n3)
	h2.Insert(n4)

	h3 := h1.Union(h2)

	if h3.min != n3 {
		t.Error("Merged heap's min should be the minimum of two heaps")
	}
}

func TestLink(t *testing.T) {
	n1 := MakeNode(1)
	n1.left = n1
	n1.right = n1
	n2 := MakeNode(3)
	n2.left = n2
	n2.right = n2

	n3 := MakeNode(5)
	n3.left = n3
	n3.right = n3

	t.Run("link to node without child", func(t *testing.T) {
		link(n1, n2)
		if n1.parent != n2 {
			t.Error("link should assign node's parent")
		}

		if n1.mark != false {
			t.Error("link should unmark the node")
		}

		if n1.left != n1 || n1.right != n1 {
			t.Error("link should create list with single element")
		}
	})
	t.Run("link to node with a child", func(t *testing.T) {
		link(n3, n2)

		if n2.child.left != n3 {
			t.Error("link should add node on the left side of parent's child")
		}
	})
}

func TestConsolidate(t *testing.T) {
	h := MakeHeap()
	n1 := MakeNode(1)
	n1.left = n1
	n1.right = n1
	h.Insert(n1)
	n2 := MakeNode(2)
	n2.left = n2
	n2.right = n2
	h.Insert(n2)

	h.consolidate()
	if h.Minimum() != n1 {
		t.Error("consolidate should update heaps mininum element")
	}

	if h.min.child != n2 {
		t.Error("consolidate should make heap's root list contain only single root with given degree")
	}

}
