package singlylinked

import "fmt"

// Node type is a node of singly linked list.
type Node struct {
	key  float32
	next *Node
}

func makeNode(v float32) *Node {
	return &Node{key: v, next: nil}
}

// GetKey is a getter for the key property of a node.
func (n *Node) GetKey() float32 {
	return n.key
}

// List is a structure holding singly linked list.
type List struct {
	head *Node
}

// MakeList creates an empty list.
func MakeList() *List {
	return &List{}
}

// IsEmpty checks if list is empty.
func (l *List) IsEmpty() bool {
	return l.head == nil
}

// GetHead returns value of listt head.
func (l *List) GetHead() *Node {
	return l.head
}

func (l *List) String() string {
	x := l.head
	if x == nil {
		return "Empty list!"
	}
	s := ""
	s += "->"
	for x != nil {
		s += fmt.Sprintf("%.2f->", x.key)
		x = x.next
	}
	return s
}

// GetTail returns the last node from the list
func (l *List) GetTail() *Node {
	x := l.head
	if l.head != nil {
		for x.next != nil {
			x = x.next
		}
	}
	return x
}

// Prepend is adding node with a given value at the head of the list.
func (l *List) Prepend(v float32) {
	n := makeNode(v)

	n.next = l.head
	l.head = n
}

// Append should add an item at the end of the list.
func (l *List) Append(v float32) {
	x := l.head
	if x == nil {
		l.Prepend(v)
	} else {
		for x.next != nil {
			x = x.next
		}
		n := makeNode(v)
		x.next = n
	}
}

// Pop is removing first element from the list and returns it.
func (l *List) Pop() *Node {
	x := l.head
	if l.head != nil {
		l.head = x.next
	}
	return x
}

// Search finds the first element with given value as its key.
func (l *List) Search(v float32) *Node {
	x := l.head
	for x != nil && x.key != v {
		x = x.next
	}
	return x
}

// Delete removes from a list first node of a given value.
// It returns true if node was deleted, false otherwise.
func (l *List) Delete(v float32) bool {
	x := l.head
	if x == nil {
		return false
	}
	if x.key == v {
		l.head = x.next
		return true
	}
	for x.next != nil {
		if x.next.key == v {
			x.next = x.next.next
			return true
		}
		x = x.next
	}
	return false
}

// Insert inserts a node with given value at given position.
// It returns true if node was inserted, false otherwise.
func (l *List) Insert(v float32, pos int) bool {
	if pos < 0 {
		return false
	}
	if pos == 0 {
		l.Prepend(v)
		return true
	}
	i := 1
	x := l.head
	for x.next != nil && i != pos {
		x = x.next
		i++
	}
	if i != pos {
		return false
	}
	n := makeNode(v)
	n.next = x.next
	x.next = n
	return true
}

// InsertAscending inserts a node with given values inside the list
// in such position that keeps list in ascending order.
func (l *List) InsertAscending(v float32) {
	if l.head == nil || l.head.key > v {
		l.Prepend(v)
		return
	}
	x := l.head
	for x.next != nil && x.next.key < v {
		x = x.next
	}
	n := makeNode(v)
	n.next = x.next
	x.next = n
}

// HasCycle checks if list contains a cycle.
func (l *List) HasCycle() bool {
	if l.head == nil {
		return false
	}

	slow := l.head
	fast := l.head.next
	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}
	return true
}

// Reverse reverses the order of the list.
func (l *List) Reverse() {
	var prev *Node
	current := l.head
	var next *Node
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

// Compare checks equality of two lists.
func (l *List) Compare(l2 *List) bool {
	if l.head == nil {
		if l2.head == nil {
			return true
		}
		return false
	}
	if l2.head == nil {
		return false
	}
	x1 := l.head
	x2 := l2.head
	for x1 != nil && x2 != nil {
		if x1.key != x2.key {
			return false
		}
		x1 = x1.next
		x2 = x2.next
	}
	if (x1 == nil && x2 != nil) || (x1 != nil && x2 == nil) {
		return false
	}
	return true
}

// IsAscending checks if list is sorted in ascending order.
func (l *List) IsAscending() bool {
	if l.head == nil || l.head.next == nil {
		return true
	}
	x := l.head
	for x.next != nil {
		if x.next.key < x.key {
			return false
		}
		x = x.next
	}
	return true
}

// RemoveDuplicatesFromAscendingList removes nodes with duplicated values.
func (l *List) RemoveDuplicatesFromAscendingList() {
	if l.head == nil || l.head.next == nil {
		return
	}
	x := l.head
	for x.next != nil {
		if x.key == x.next.key {
			x.next = x.next.next
		}
		if x.next != nil {
			x = x.next
		}
	}
}

// SortAscending should sort singly linked list in ascending order.
func (l *List) SortAscending() {
	l.head = l.head.sortAscending()
}

func (n *Node) sortAscending() *Node {
	if n == nil || n.next == nil {
		return n
	}
	//List *middle, *middleNext
	middle := n.findMiddle()
	middleNext := middle.next
	middle.next = nil

	//List *left, *right, *head
	left := n.sortAscending()
	right := middleNext.sortAscending()
	head := left.mergeAscendingLists(right)
	return head
}

func (n *Node) findMiddle() *Node {
	if n == nil {
		return nil
	}

	slow := n
	fast := n

	for fast != nil {
		fast = fast.next
		if fast != nil {
			fast = fast.next
		}
		if fast != nil {
			slow = slow.next
		}
	}
	return slow
}

func (n *Node) mergeAscendingLists(headB *Node) *Node {
	if n == nil {
		return headB
	}
	if headB == nil {
		return n
	}
	var head, h *Node
	for n != nil && headB != nil {
		if n.key > headB.key {
			if h == nil {
				h = headB
				head = headB
			} else {
				h.next = headB
				h = h.next
			}
			headB = headB.next
		} else {
			if h == nil {
				h = n
				head = n
			} else {
				h.next = n
				h = h.next
			}
			n = n.next
		}
	}
	if n != nil {
		h.next = n
	} else {
		h.next = headB
	}
	return head
}
