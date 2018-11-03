package singlylinked

import (
	"fmt"
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
				t.Errorf("makeNode(%d): expected: %d got %d", test.key, got.key, test.expected)
			}

			if got.next != nil {
				t.Error("makeNode should produce node with nil as its next property")
			}
		})
	}
}

func TestMakeList(t *testing.T) {
	l := MakeList()
	if l.head != nil {
		t.Error("MakeList should create an empty list")
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
	if l.Delete(3) {
		t.Error("Delete should return false when called on empty list")
	}
	l.Prepend(3)
	l.Prepend(4)
	l.Prepend(5)

	if !l.Delete(3) {
		t.Error("Delete should return true when removing an element")
	}
	if l.head.next.next != nil {
		t.Error("Delete should remove node from a list")
	}
	if l.head.key != 5 {
		t.Error("Delete should not change rest of the elements")
	}

	l.Delete(5)
	if l.head.key != 4 {
		t.Error("Delete should remove first element from the list")
	}

	l.Delete(4)

	if l.head != nil {
		t.Error("Delete should remove last element from the list")
	}
}

func TestAppend(t *testing.T) {
	l := MakeList()
	l.Append(4)

	if l.head.key != 4 {
		t.Error("Append should at the first item to the list")
	}

	l.Append(5)
	if l.head.next.key != 5 {
		t.Error("Append should add a node at the end of the list")
	}

	if l.head.key != 4 {
		t.Error("Append should not update list head")
	}
}

func ExamplePrint() {
	l := MakeList()
	l.Print()
	l.Prepend(3)
	l.Print()
	l.Prepend(4)
	l.Print()
	l.Prepend(5)
	l.Print()

	// Output:
	// Empty list!
	// ->3->
	// ->4->3->
	// ->5->4->3->
}

func ExampleInsert() {
	l := MakeList()
	l.Insert(4, 0)
	l.Print()
	fmt.Println(l.Insert(4, 2))
	l.Print()
	l.Insert(5, 1)
	l.Print()
	l.Insert(6, 2)
	l.Print()
	l.Insert(3, 1)
	l.Print()
	l.Insert(8, 0)
	l.Print()

	// Output:
	// ->4->
	// false
	// ->4->
	// ->4->5->
	// ->4->5->6->
	// ->4->3->5->6->
	// ->8->4->3->5->6->
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
	l.Print()
	l.Reverse()
	l.Print()

	// Output:
	// ->2->7->5->4->
	// ->4->5->7->2->
}

func ExampleCompare() {
	l := MakeList()
	l2 := MakeList()
	fmt.Println(l.Compare(l2))
	l.Append(4)
	fmt.Println(l.Compare(l2))
	l2.Append(4)
	fmt.Println(l.Compare(l2))
	l.Append(5)
	fmt.Println(l.Compare(l2))
	l2.Append(5)
	fmt.Println(l.Compare(l2))
	l2.Append(7)
	l.Append(8)
	fmt.Println(l.Compare(l2))

	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
}

func ExampleInsertSort() {
	l := MakeList()

	l.InsertSort(4)
	l.Print()
	l.InsertSort(7)
	l.Print()
	l.InsertSort(3)
	l.Print()
	l.InsertSort(8)
	l.Print()
	l.InsertSort(0)
	l.Print()
	l.InsertSort(5)
	l.Print()

	// Output:
	// ->4->
	// ->4->7->
	// ->3->4->7->
	// ->3->4->7->8->
	// ->0->3->4->7->8->
	// ->0->3->4->5->7->8->
}

func ExampleRemoveDuplicatesFromAscendingList() {
	l := MakeList()

	l.RemoveDuplicatesFromAscendingList()
	l.Print()
	l.InsertSort(3)
	l.RemoveDuplicatesFromAscendingList()
	l.Print()
	l.InsertSort(3)
	l.Print()
	l.RemoveDuplicatesFromAscendingList()
	l.Print()
	l.InsertSort(7)
	l.RemoveDuplicatesFromAscendingList()
	l.Print()
	l.InsertSort(3)
	l.Print()
	l.RemoveDuplicatesFromAscendingList()
	l.Print()

	// Output:
	// Empty list!
	// ->3->
	// ->3->3->
	// ->3->
	// ->3->7->
	// ->3->3->7->
	// ->3->7->
}

func ExampleIsAscending() {
	l := MakeList()
	fmt.Println(l.IsAscending())
	l.Prepend(2)
	fmt.Println(l.IsAscending())
	l.Prepend(4)
	fmt.Println(l.IsAscending())
	l.Prepend(1)
	fmt.Println(l.IsAscending())
	l.Delete(4)
	fmt.Println(l.IsAscending())
	l.Append(5)
	fmt.Println(l.IsAscending())

	// Output:
	// true
	// true
	// false
	// false
	// true
	// true
}

func TestGetKey(t *testing.T) {
	var testsTable = []struct {
		name     string
		node     Node
		expected int
	}{
		{"positive number #1", Node{4, nil}, 4},
		{"positive number #2", Node{32321, nil}, 32321},
		{"negative number", Node{-21, nil}, -21},
		{"empty node", Node{}, 0},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			got := test.node.GetKey()
			if got != test.expected {
				t.Errorf("%v.GetKey(): expected: %d got: %d", test.node, got, test.expected)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	var testsTable = []struct {
		name     string
		node     *Node
		expected bool
	}{
		{"empty list", nil, true},
		{"valid node", &Node{32321, nil}, false},
		{"empty node", &Node{}, false},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			l := MakeList()
			l.head = test.node
			got := l.IsEmpty()
			if got != test.expected {
				t.Errorf("l.IsEmpty(): expected: %t got: %t", got, test.expected)
			}
		})
	}
}

func TestGetHead(t *testing.T) {
	var testsTable = []struct {
		name string
		node *Node
	}{
		{"empty list", nil},
		{"valid node", &Node{32321, nil}},
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
