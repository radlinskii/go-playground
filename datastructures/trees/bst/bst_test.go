package bst

import (
	"testing"
)

func TestMakeNode(t *testing.T) {
	var testsTable = []struct {
		name     string
		key      int
		expected int
	}{
		{"negative number", -321, -321},
		{"zero", 0, 0},
		{"positive number", 321, 321},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := makeNode(test.key)
			if got.key != test.expected {
				t.Errorf("makeNode(%d): expected: %d, got %d", test.key, got.key, test.expected)
			}

			if got.left != nil {
				t.Error("makeNode should produce node with nil as its left property")
			}
			if got.right != nil {
				t.Error("makeNode should produce node with nil as its right property")
			}
			if got.parent != nil {
				t.Error("makeNode should produce node with nil as its parent property")
			}
		})
	}
}

func TestMakeTree(t *testing.T) {
	tree := MakeTree()
	if tree.root != nil {
		t.Error("MakeTree should create an empty tree")
	}
}

func TestGetKey(t *testing.T) {
	var testsTable = []struct {
		name     string
		node     Node
		expected int
	}{
		{"negative key", Node{key: -321}, -321},
		{"undefined key", Node{}, 0},
		{"positive key", Node{key: 321}, 321},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := test.node.GetKey()
			if got != test.expected {
				t.Errorf("makeNode(%v): expected: %d, got %d", test.node, got, test.expected)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tree := MakeTree()

	n5 := tree.Insert(5)
	if tree.root != n5 {
		t.Error("Insert to empty tree should update tree's root ")
	}
	if tree.root.parent != nil {
		t.Error("Insert to empty tree should leave root's parent as nil")
	}

	n3 := tree.Insert(3)
	if n5.left != n3 {
		t.Error("Insert lower value than the root should update root's left child")
	}

	n4 := tree.Insert(4)
	if n3.right != n4 {
		t.Error("Insert should always add node with greater values as the right child")
	}
}

func TestDelete(t *testing.T) {
	tree := MakeTree()
	n6 := tree.Insert(6)
	n3 := tree.Insert(3)
	n5 := tree.Insert(5)
	n4 := tree.Insert(4)
	n2 := tree.Insert(2)

	tree.Delete(n3)
	if n4 != n6.left {
		t.Error("Delete node with two children should replace node with the lowest value from its right subtree")
	}

	tree.Delete(n5)
	if n4.right != nil {
		t.Error("Delete node without children should replace it with nil")
	}

	tree.Delete(n4)
	if n6.left != n2 {
		t.Error("Removing node with one child should replace it with its child")
	}
}

func TestTransplant(t *testing.T) {
	tree := MakeTree()

	n1 := tree.Insert(1)

	n2 := makeNode(2)

	tree.transplant(n1, n2)
	if tree.root != n2 {
		t.Error("transplant should replace root properly")
	}
	tree2 := MakeTree()
	n4 := tree2.Insert(4)
	n3 := tree2.Insert(3)
	n5 := tree2.Insert(5)

	n9 := makeNode(9)
	n10 := makeNode(10)

	tree2.transplant(n5, n9)
	if n4.right != n9 {
		t.Error("transplant should replace node when its a right child")
	}
	if n9.parent != n4 {
		t.Error("transplant should update new node's parent")
	}

	tree2.transplant(n3, n10)
	if n4.left != n10 {
		t.Error("transplant should replace node when it's a left child")
	}
}

func TestSearch(t *testing.T) {
	tree := MakeTree()
	n4 := tree.Insert(4)
	n2 := tree.Insert(2)
	n5 := tree.Insert(5)

	if n4 != tree.Search(4) {
		t.Error("Search should find root value")
	}
	if n2 != tree.Search(2) {
		t.Error("Search should find node on the left side")
	}
	if n5 != tree.Search(5) {
		t.Error("Search should find node on the right side")
	}

	if tree.Search(6) != nil {
		t.Error("Search should not find unexisting node")
	}
}

func TestSearchRecursive(t *testing.T) {
	tree := MakeTree()
	n4 := tree.Insert(4)
	n2 := tree.Insert(2)
	n5 := tree.Insert(5)

	if n4 != tree.SearchRecursive(4) {
		t.Error("SearchRecursive should find root value")
	}
	if n2 != tree.SearchRecursive(2) {
		t.Error("SearchRecursive should find node on the left side")
	}
	if n5 != tree.SearchRecursive(5) {
		t.Error("SearchRecursive should find node on the right side")
	}

	if tree.SearchRecursive(6) != nil {
		t.Error("SearchRecursive should not find unexisting node")
	}
}

func TestMinimum(t *testing.T) {
	tree := MakeTree()
	tree.Insert(4)
	tree.Insert(2)
	n1 := tree.Insert(1)

	if n1 != tree.Minimum() {
		t.Error("Minimum should return the node of minimum key")
	}
}

func TestMinimumRecursive(t *testing.T) {
	tree := MakeTree()
	tree.Insert(4)
	tree.Insert(2)
	n1 := tree.Insert(1)

	if n1 != tree.MinimumRecursive() {
		t.Error("MinimumRecursive should return the node of minimum key")
	}
}

func TestMaximum(t *testing.T) {
	tree := MakeTree()
	tree.Insert(1)
	tree.Insert(2)
	n4 := tree.Insert(4)

	if n4 != tree.Maximum() {
		t.Error("Maximum should return the node of maximum key")
	}
}

func TestMaximumRecursive(t *testing.T) {
	tree := MakeTree()
	tree.Insert(1)
	tree.Insert(2)
	n4 := tree.Insert(4)

	if n4 != tree.MaximumRecursive() {
		t.Error("MaximumRecursive should return the node of maximum key")
	}
}

func TestSuccessor(t *testing.T) {
	tree := MakeTree()
	n5 := tree.Insert(5)
	tree.Insert(7)
	tree.Insert(8)
	n6 := tree.Insert(6)
	tree.Insert(3)
	n4 := tree.Insert(4)

	if n5.Successor() != n6 {
		t.Error("Successor should return min from non-nil right subtree")
	}
	if n4.Successor() != n5 {
		t.Error("Successor should search up when right subtree is nil")
	}
}

func TestPredecessor(t *testing.T) {
	tree := MakeTree()
	n5 := tree.Insert(5)
	tree.Insert(8)
	n6 := tree.Insert(6)
	n4 := tree.Insert(4)

	if n5.Predecessor() != n4 {
		t.Error("Predecessor should return max from non-nil left subtree")
	}
	if n6.Predecessor() != n5 {
		t.Error("Predecessor should search up when left subtree is nil")
	}
}

func ExampleInOrder() {
	tree := MakeTree()
	tree.Insert(6)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(7)

	tree.InOrder()

	// Output:
	// 1
	// 3
	// 4
	// 6
	// 7
	// 8
}

func ExamplePreOrder() {
	tree := MakeTree()
	tree.Insert(6)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(7)

	tree.PreOrder()

	// Output:
	// 6
	// 3
	// 1
	// 4
	// 8
	// 7
}

func ExamplePostOrder() {
	tree := MakeTree()
	tree.Insert(6)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(7)

	tree.PostOrder()

	// Output:
	// 1
	// 4
	// 3
	// 7
	// 8
	// 6
}
