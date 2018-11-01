package singlylinked

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

// Insert is adding node with a given value at the head of the list.
func (l *List) Insert(v int) {
	n := &Node{key: v}

	n.next = l.head
	l.head = n
}

// Search finds the first element with given value as its key.
func (l *List) Search(v int) *Node {
	x := l.head
	for x != nil && x.key != v {
		x = x.next
	}
	return x
}
