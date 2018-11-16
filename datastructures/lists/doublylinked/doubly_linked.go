package doublylinked

import "fmt"

// Node type is a node of doubly linked list.
type Node struct {
	next *Node
	prev *Node
	key  float32
}

func makeNode(v float32) *Node {
	return &Node{key: v}
}

// GetKey is a getter for the key property of a node.
func (n *Node) GetKey() float32 {
	return n.key
}

// List is a structure holding doubly linked list.
type List struct {
	head *Node
	tail *Node
}

// MakeList creates an empty list.
func MakeList() *List {
	return &List{}
}

// IsEmpty checks if list is empty.
func (l *List) IsEmpty() bool {
	return l.head == nil && l.tail == nil
}

// GetHead returns value of list head.
func (l *List) GetHead() *Node {
	return l.head
}

// GetTail returns value of list tail.
func (l *List) GetTail() *Node {
	return l.tail
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

// Prepend is adding node with a given value at the head of the list.
func (l *List) Prepend(v float32) {
	n := makeNode(v)

	if l.IsEmpty() {
		l.tail = n
	} else {
		n.next = l.head
		l.head.prev = n
	}
	l.head = n
}

// Append should add an item at the end of the list.
func (l *List) Append(v float32) {
	if l.IsEmpty() {
		l.Prepend(v)
	} else {
		n := makeNode(v)
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
}

// Pop is removing first element from the list and returns it.
func (l *List) Pop() *Node {
	x := l.head
	if !l.IsEmpty() {
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

// Delete removes given node from the list.
func (l *List) Delete(n *Node) {
	if n.prev == nil {
		l.head = n.next
	} else {
		n.prev.next = n.next
	}
	if n.next == nil {
		l.tail = n.prev
	} else {
		n.next.prev = n.prev
	}
}

// HasCycle checks if there is a cycle in the list.
func (l *List) HasCycle() bool {
	if l.head == nil {
		return false
	}

	slow := l.head
	fast := l.head.next
	for fast != slow {
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
	var tmp *Node
	curr := l.head
	for curr != nil {
		tmp = curr.next
		curr.prev, curr.next = curr.next, curr.prev
		curr = tmp
	}
	l.head, l.tail = l.tail, l.head
}
