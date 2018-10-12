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
	Key                        int
	left, right, parent, child *Node
	mark                       bool
	degree                     int
}

func (fh *Heap) addRoot(x *Node) {
	if fh.min == nil {
		// create fh's root list containing only x
		x.left = x
		x.right = x
		fh.min = x
	} else {
		// insert x to fh's root list
		fh.min.left.right = x
		x.right = fh.min
		x.left = fh.min.left
		fh.min.left = x
		if x.Key < fh.min.Key {
			fh.min = x
		}
	}
}

// MakeHeap creates and returns a new, empty heap.
func MakeHeap() *Heap {
	return &Heap{}
}

// Insert inserts a new node, with predeclared Key, to the heap.
func (fh *Heap) Insert(x *Node) *Node {
	x.degree = 0
	x.mark = false
	x.parent = nil
	x.child = nil

	fh.addRoot(x)
	fh.n++
	return x
}

// Minimum returns pointer to the heap's node holding the minimum Key.
func (fh *Heap) Minimum() *Node {
	return fh.min
}

// Union creates and returns the merge of two mergeable heaps.
func (fh *Heap) Union(fh2 *Heap) *Heap {
	newFH := MakeHeap()
	newFH.min = fh.min

	newFH.min.left.right = fh2.min
	fh2.min.left.right = newFH.min
	fh2.min.left, newFH.min.left = newFH.min.left, fh2.min.left

	if fh.min == nil || (fh2.min != nil && fh.min.Key > fh2.min.Key) {
		newFH.min = fh2.min
	}
	newFH.n = fh.n + fh2.n
	return newFH
}

// ExtractMin extracts the node with minimum Key from a heap
// and returns pointer to this node.
func (fh *Heap) ExtractMin() *Node {
	z := fh.min
	if z != nil {
		for {
			// add z children to fh's root list
			if x := z.child; x != nil {
				x.parent = nil
				if x.right != x {
					z.child = x.right
					x.right.left = x.left
					x.left.right = x.right
				} else {
					z.child = nil
				}
				x.left = z.left
				x.right = z
				z.left.right = x
				z.left = x
			} else {
				break
			}
		}
		// remove z from fh's root list
		z.left.right = z.right
		z.right.left = z.left

		if z == z.right {
			fh.min = nil
		} else {
			fh.min = z.right
			fh.consolidate()
		}
		fh.n--
	}
	return z
}

func (fh *Heap) consolidate() {
	degreeToRoot := make(map[int]*Node)
	w := fh.min
	last := w.left
	for {
		r := w.right
		x := w
		d := x.degree
		for {
			if y, ok := degreeToRoot[d]; !ok {
				break
			} else {
				if y.Key < x.Key {
					y, x = x, y
				}
				fh.link(y, x)
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
	fh.min = nil
	for _, v := range degreeToRoot {
		fh.addRoot(v)
	}

}

func (fh *Heap) link(y, x *Node) {
	// remove y from fh's root list
	y.right.left = y.left
	y.left.right = y.right
	// make y a child of x and increase degree of x
	y.parent = x
	if x.child == nil {
		x.child = y
		y.left = y
		y.right = y
	} else {
		y.left = x.child.left
		y.right = x.child
		x.child.left.right = y
		x.child.left = y
	}

	y.mark = false
}

// DecreaseKey decreases the key of given node.
func (fh *Heap) DecreaseKey(x *Node, k int) {
	if x.Key < k {
		panic("new Key is greater than the previous one")
	}
	x.Key = k
	y := x.parent
	if y != nil && x.Key < y.Key {
		fh.cut(x, y)
		fh.cascadingCut(y)
	}
	if x.Key < fh.min.Key {
		fh.min = x
	}
}

func (fh *Heap) cut(x, y *Node) {
	// remove x from y's children list and decrement y's degree
	if x.right != x {
		y.child = x.right
		x.right.left = x.left
		x.left.right = x.right
	} else {
		y.child = nil
	}
	y.degree--
	// add x to fh's root list
	x.left = fh.min.left
	x.right = fh.min
	fh.min.left.right = x
	fh.min.left = x

	x.parent = nil
	x.mark = false
}

func (fh *Heap) cascadingCut(y *Node) {
	z := y.parent
	if z != nil {
		if !y.mark {
			y.mark = true
		} else {
			fh.cut(y, z)
			fh.cascadingCut(z)
		}
	}
}

// Delete deletes node x from heap fh.
func (fh *Heap) Delete(x *Node) {
	fh.DecreaseKey(x, int(-1<<63))
	fh.ExtractMin()
}

// Vis visualizes the heap. All credits to "https://rosettacode.org/wiki/Fibonacci_heap"
func (fh Heap) Vis() {
	if fh.min == nil {
		fmt.Println("<empty>")
		return
	}
	var f func(*Node, string)
	f = func(n *Node, pre string) {
		pc := "│ "
		for x := n; ; x = x.right {
			if x.right != n {
				fmt.Print(pre, "├─")
			} else {
				fmt.Print(pre, "└─")
				pc = "  "
			}
			if x.child == nil {
				fmt.Println("╴", x.Key)
			} else {
				fmt.Println("┐", x.Key)
				f(x.child, pre+pc)
			}
			if x.right == n {
				break
			}
		}
	}
	f(fh.min, "")
}
