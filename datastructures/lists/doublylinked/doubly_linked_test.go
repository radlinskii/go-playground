package doublylinked

import (
	"fmt"
	"testing"
)

func TestMakeNode(t *testing.T) {
	var testsTable = []struct {
		name     string
		key      float32
		expected float32
	}{
		{"negative number", -321, -321},
		{"zero", 0, 0},
		{"positive number", 321, 321},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := makeNode(test.key)
			if got.key != test.expected {
				t.Errorf("makeNode(%f): expected: %f got %f", test.key, got.key, test.expected)
			}

			if got.next != nil {
				t.Error("makeNode should produce node with nil as its next property")
			}
			if got.prev != nil {
				t.Error("makeNode should produce node with nil as its prev property")
			}
		})
	}
}

func TestGetKey(t *testing.T) {
	var testsTable = []struct {
		name     string
		node     Node
		expected float32
	}{
		{"positive number #1", Node{key: 4}, 4},
		{"positive number #2", Node{key: 32321}, 32321},
		{"negative number", Node{key: -21}, -21},
		{"empty node", Node{}, 0},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := test.node.GetKey()
			if got != test.expected {
				t.Errorf("%v.GetKey(): expected: %f got: %f", test.node, got, test.expected)
			}
		})
	}
}

func TestMakeList(t *testing.T) {
	l := MakeList()
	if l.head != nil {
		t.Error("MakeList should create a list with empty head")
	}
	if l.tail != nil {
		t.Error("MakeList should create a list with empty tail")
	}
}

func TestGetHead(t *testing.T) {
	var testsTable = []struct {
		name string
		node *Node
	}{
		{"empty list", nil},
		{"valid node", &Node{key: 32321}},
		{"empty node", &Node{}},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			l := MakeList()
			l.head = test.node
			got := l.GetHead()
			if got != test.node {
				t.Errorf("l.GetHead(): expected: %v got: %v", test.node, got)
			}
		})
	}
}
func TestGetTail(t *testing.T) {
	var testsTable = []struct {
		name string
		node *Node
	}{
		{"empty list", nil},
		{"valid node", &Node{key: 32321}},
		{"empty node", &Node{}},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			l := MakeList()
			l.tail = test.node
			got := l.GetTail()
			if got != test.node {
				t.Errorf("l.GetHead(): expected: %v got: %v", test.node, got)
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	l := MakeList()
	l.Prepend(3)

	if l.head.key != 3 {
		t.Error("Prepend should insert 3 as key of list head")
	}

	l.Prepend(5)
	if l.head.key != 5 {
		t.Error("Prepend should insert 5 as key of list head")
	}

	if l.head.next.key != 3 {
		t.Error("Prepend add new node at the begging of the list")
	}
}

func TestAppend(t *testing.T) {
	l := MakeList()
	l.Append(4)

	if l.head.key != 4 {
		t.Error("Append should at the first item to the list as its head")
	}

	if l.tail.key != 4 {
		t.Error("Append should at the first item to the list as its tail")
	}

	l.Append(5)
	if l.head.next.key != 5 {
		t.Error("Append should add a node at the end of the list #1")
	}

	if l.tail.key != 5 {
		t.Error("Append should add a node at the end of the list #2")
	}

	if l.head.key != 4 {
		t.Error("Append should not update list head")
	}
}

func TestString(t *testing.T) {
	l := MakeList()
	if l.String() != "Empty list!" {
		t.Error("String on empty list failed")
		t.Error(l)
	}

	l.Prepend(5)
	l.Prepend(7)
	l.Prepend(6)
	l.Prepend(4)

	if l.String() != "->4.00->6.00->7.00->5.00->" {
		t.Error("String failure")
		t.Error(l)
	}
}

func TestPop(t *testing.T) {
	l := MakeList()
	n := l.Pop()
	if n != nil {
		t.Error("Pop should return nil when called on an empty list")
	}

	l.Prepend(4)

	n = l.Pop()
	if n.key != 4 {
		t.Error("Pop should return first element from a list")
	}

	if l.head != nil {
		t.Error("Pop should remove an element from a list")
	}
}

func TestSearch(t *testing.T) {
	l := MakeList()

	if l.Search(3) != nil {
		t.Error("Search on empty list should definitely return nil")
	}

	l.Prepend(3)

	if l.Search(3) != l.head {
		t.Error("Search should find the right node in list with single element")
	}

	l.Prepend(4)
	if l.Search(3) != l.head.next {
		t.Error("Search should find nodes deep in the list")
	}

	if l.Search(5) != nil {
		t.Error("Search should not find unexisting values")
	}
}

func TestDelete(t *testing.T) {
	l := MakeList()
	l.Prepend(3)
	l.Prepend(4)
	l.Prepend(5)

	n := l.Search(3)
	l.Delete(n)
	if l.head.next.next != nil {
		t.Error("Delete should remove node from a list")
	}
	if l.head.key != 5 {
		t.Error("Delete should not change rest of the elements")
	}

	n = l.Search(5)
	l.Delete(n)
	if l.head.key != 4 {
		t.Error("Delete should remove first element from the list")
	}

	n = l.Search(4)
	l.Delete(n)
	if l.head != nil {
		t.Error("Delete should remove last element from the list")
	}
}

func TestHasCycle1(t *testing.T) {
	l := MakeList()
	if l.HasCycle() {
		t.Error("HasCycle shouldn't detect a cycle without a reason")
	}
	l.Prepend(2)
	l.Prepend(3)
	l.Prepend(4)

	if l.HasCycle() {
		t.Error("HasCycle shouldn't detect a cycle without a reason")
	}
}

func TestHasCycle2(t *testing.T) {
	l := MakeList()

	l.Prepend(3)
	l.Prepend(4)
	l.Prepend(5)
	n := l.Search(3)
	n.next = l.head

	if !l.HasCycle() {
		t.Error("HasCycle should properly detect a cycle")
	}
}

func ExampleReverse() {
	l := MakeList()
	l.Prepend(4)
	l.Prepend(5)
	l.Prepend(7)
	l.Prepend(2)
	fmt.Println(l)
	l.Reverse()
	fmt.Println(l)

	// Output:
	// ->2.00->7.00->5.00->4.00->
	// ->4.00->5.00->7.00->2.00->
}
