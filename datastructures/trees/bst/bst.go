package trees

import "fmt"

// Node is a Node of Binary Search Tree.
type Node struct {
	key                 int
	left, right, parent *Node
}

// GetKey returns key of a node.
func (n *Node) GetKey() int {
	return n.key
}

func makeNode(v int) *Node {
	return &Node{key: v}
}

// Tree is a data structure representing Binary Serach Tree.
type Tree struct {
	root *Node
}

// MakeTree returns pointer to a new Binary Search Tree.
func MakeTree() *Tree {
	return &Tree{}
}

// InOrder walks the tree in the sorted order.
func (t *Tree) InOrder() {
	t.root.inOrder()
}

func (n *Node) inOrder() {
	if n != nil {
		n.left.inOrder()
		fmt.Println(n.key)
		n.right.inOrder()
	}
}

// PreOrder walks the tree in the preorder style.
func (t *Tree) PreOrder() {
	t.root.preOrder()
}

func (n *Node) preOrder() {
	if n != nil {
		fmt.Println(n.key)
		n.left.preOrder()
		n.right.preOrder()
	}
}

// PostOrder walks the tree in the postorder style.
func (t *Tree) PostOrder() {
	t.root.postOrder()
}

func (n *Node) postOrder() {
	if n != nil {
		n.left.postOrder()
		n.right.postOrder()
		fmt.Println(n.key)
	}
}

// Search finds node with a given key in the tree.
func (t *Tree) Search(v int) *Node {
	return t.root.search(v)
}

func (n *Node) search(v int) *Node {
	x := n
	for x != nil && v != x.key {
		if v < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

// SearchRecursive finds node with a given key in the tree. Recursively.
func (t *Tree) SearchRecursive(v int) *Node {
	return t.root.searchRecursive(v)
}

func (n *Node) searchRecursive(v int) *Node {
	if n == nil || v == n.key {
		return n
	}
	if v < n.key {
		return n.left.searchRecursive(v)
	}
	return n.right.searchRecursive(v)
}

// Minimum finds node with the smallest key stored in the tree.
func (t *Tree) Minimum() *Node {
	return t.root.minimum()
}

func (n *Node) minimum() *Node {
	x := n
	if x != nil {
		for x.left != nil {
			x = x.left
		}
	}
	return x
}

// MinimumRecursive finds node with the smallest key stored in the tree. Recursively.
func (t *Tree) MinimumRecursive() *Node {
	return t.root.minimumRecursive()
}

func (n *Node) minimumRecursive() *Node {
	if n == nil || n.left == nil {
		return n
	}
	return n.left.minimumRecursive()
}

// Maximum finds node with the largest key stored in the tree.
func (t *Tree) Maximum() *Node {
	return t.root.maximum()
}

func (n *Node) maximum() *Node {
	x := n
	if x != nil {
		for x.right != nil {
			x = x.right
		}
	}
	return x
}

// MaximumRecursive finds node with the largest key stored in the tree. Recursively.
func (t *Tree) MaximumRecursive() *Node {
	return t.root.maximumRecursive()
}

func (n *Node) maximumRecursive() *Node {
	if n == nil || n.right == nil {
		return n
	}
	return n.right.maximumRecursive()
}

// Successor returns node with the smallest key greater than the given node's key.
func (n *Node) Successor() *Node {
	if n.right != nil {
		return n.right.minimum()
	}
	x := n
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

// Predecessor returns the node with the greatest key lower than the given node's key.
func (n *Node) Predecessor() *Node {
	if n.left != nil {
		return n.left.maximum()
	}
	x := n
	y := x.parent
	for y != nil && x == y.left {
		x = y
		y = y.parent
	}
	return y
}

// Insert inserts node with a given value to the tree.
func (t *Tree) Insert(v int) *Node {
	z := makeNode(v)
	x := t.root
	var y *Node
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	return z
}

// Delete removes a given node from the tree.
func (t *Tree) Delete(z *Node) {
	if z.left == nil {
		t.transplant(z, z.right)
	} else if z.right == nil {
		t.transplant(z, z.left)
	} else {
		y := z.right.minimum()
		if y.parent != z {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
	}
}

func (t *Tree) transplant(u, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}
