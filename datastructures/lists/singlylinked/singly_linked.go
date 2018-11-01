package singlylinked

import "fmt"

// Node type is a node of singly linked list.
type Node struct {
	key  int
	next *Node
}

func makeNode(v int) *Node {
	return &Node{key: v, next: nil}
}

// List is a structure holding singly linked list.
type List struct {
	head *Node
}

// MakeList creates an empty list.
func MakeList() *List {
	return &List{}
}

// Print prints keys of elements of the list.
func (l *List) Print() {
	x := l.head
	if x == nil {
		fmt.Println("Empty list!")
	} else {
		fmt.Print("->")
		for x != nil {
			fmt.Printf("%d->", x.key)
			x = x.next
		}
		fmt.Println()
	}
}

// Prepend is adding node with a given value at the head of the list.
func (l *List) Prepend(v int) {
	n := &Node{key: v}

	n.next = l.head
	l.head = n
}

// Append should add an item at the end of the list
func (l *List) Append(v int) {
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
func (l *List) Search(v int) *Node {
	x := l.head
	for x != nil && x.key != v {
		x = x.next
	}
	return x
}

// Delete removes from a list first node of a given value.
// It returns true if node was deleted, false otherwise.
func (l *List) Delete(v int) bool {
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
func (l *List) Insert(v, pos int) bool {
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
	n := &Node{key: v}
	n.next = x.next
	x.next = n
	return true
}
