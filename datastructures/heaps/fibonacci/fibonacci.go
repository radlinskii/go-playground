package fibonacci

import "fmt"

// Heap is a implementation of Fibonacci heap.
// Implementation from Introduction to Algorithms by T. Cormen
// Reference: https://en.wikipedia.org/wiki/Fibonacci_heap
type Heap struct {
	min *Node
	n   int
}

// Node holds structure of nodes inside Fibonacci heap.
type Node struct {
	key                        int
	left, right, parent, child *Node
	mark                       bool
	degree                     int
}

// GetKey gets a value of node's key.
func (n *Node) GetKey() int {
	return n.key
}

// MakeNode creates a node with key equal to the given value.
func MakeNode(k int) *Node {
	return &Node{key: k}
}

func addNode(n1, n2 *Node) {
	n1.left.right = n2
	n2.right = n1
	n2.left = n1.left
	n1.left = n2
}

func (h *Heap) addRoot(x *Node) {
	if h.min == nil {
		// create h's root list containing only x
		x.left = x
		x.right = x
		h.min = x
	} else {
		// insert x to h's root list
		addNode(h.min, x)
		if x.key < h.min.key {
			h.min = x
		}
	}
}

func removeNodeFromList(n *Node) {
	n.left.right = n.right
	n.right.left = n.left
}

// MakeHeap creates and returns a new, empty heap.
func MakeHeap() *Heap {
	return &Heap{}
}

// Insert inserts a new node, with predeclared key, to the heap.
func (h *Heap) Insert(x *Node) *Node {
	x.degree = 0
	x.mark = false
	x.parent = nil
	x.child = nil

	h.addRoot(x)
	h.n++
	return x
}

// Minimum returns pointer to the heap's node holding the minimum key.
func (h *Heap) Minimum() *Node {
	return h.min
}

// Union creates and returns the merge of two mergeable heaps.
func (h *Heap) Union(fh2 *Heap) *Heap {
	newFH := MakeHeap()
	newFH.min = h.min

	newFH.min.left.right = fh2.min
	fh2.min.left.right = newFH.min
	fh2.min.left, newFH.min.left = newFH.min.left, fh2.min.left

	if h.min == nil || (fh2.min != nil && h.min.key > fh2.min.key) {
		newFH.min = fh2.min
	}
	newFH.n = h.n + fh2.n
	return newFH
}

// ExtractMin extracts the node with minimum key from a heap
// and returns pointer to this node.
func (h *Heap) ExtractMin() *Node {
	z := h.min
	if z != nil {
		for {
			// add z children to h's root list
			if x := z.child; x != nil {
				x.parent = nil
				if x.right != x {
					z.child = x.right
					removeNodeFromList(x)
				} else {
					z.child = nil
				}
				addNode(z, x)
			} else {
				break
			}
		}
		removeNodeFromList(z)

		if z == z.right {
			h.min = nil
		} else {
			h.min = z.right
			h.consolidate()
		}
		h.n--
	}
	return z
}

func (h *Heap) consolidate() {
	degreeToRoot := make(map[int]*Node)
	w := h.min
	last := w.left
	for {
		r := w.right
		x := w
		d := x.degree
		for {
			if y, ok := degreeToRoot[d]; !ok {
				break
			} else {
				if y.key < x.key {
					y, x = x, y
				}
				link(y, x)
				delete(degreeToRoot, d)
				d++
			}
		}
		degreeToRoot[d] = x
		if w == last {
			break
		}
		w = r
	}
	h.min = nil
	for _, v := range degreeToRoot {
		h.addRoot(v)
	}
}

func link(y, x *Node) {
	removeNodeFromList(y)
	// make y a child of x and increase degree of x
	y.parent = x
	if x.child == nil {
		x.child = y
		y.left = y
		y.right = y
	} else {
		addNode(x.child, y)
	}

	y.mark = false
}

type errInvalidArgument int

func (e errInvalidArgument) Error() string {
	return fmt.Sprintf("Cannot decrese key to a bigger value: %d", int(e))
}

// DecreaseKey decreases the key of given node.
func (h *Heap) DecreaseKey(x *Node, k int) error {
	if x.key < k {
		return errInvalidArgument(k)
	}
	x.key = k
	y := x.parent
	if y != nil && x.key < y.key {
		h.cut(x, y)
		h.cascadingCut(y)
	}
	if x.key < h.min.key {
		h.min = x
	}
	return nil
}

func (h *Heap) cut(x, y *Node) {
	// remove x from y's children list and decrement y's degree
	if x.right != x {
		y.child = x.right
		removeNodeFromList(x)
	} else {
		y.child = nil
	}
	y.degree--
	addNode(h.min, x)

	x.parent = nil
	x.mark = false
}

func (h *Heap) cascadingCut(y *Node) {
	z := y.parent
	if z != nil {
		if !y.mark {
			y.mark = true
		} else {
			h.cut(y, z)
			h.cascadingCut(z)
		}
	}
}

// Delete deletes node x from heap h.
func (h *Heap) Delete(x *Node) error {
	err := h.DecreaseKey(x, int(-1<<63))
	if err != nil {
		return err
	}
	h.ExtractMin()
	return nil
}
