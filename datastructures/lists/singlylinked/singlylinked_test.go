package singlylinked

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

	if l.Delete(6) {
		t.Error("Delete should return false when given unexisting value")
	}

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

func TestInsert(t *testing.T) {
	l := MakeList()
	if l.Insert(5, -2) {
		t.Error("Insert should return false when given negative index")
	}

	if l.GetHead() != nil {
		t.Error("Insert on negative index should not add node to the list")
	}

	l.Insert(4, 0)

	l2 := MakeList()
	l2.Prepend(4)

	if !l.Compare(l2) {
		t.Error("Inserting at index 0 should work like prepending")
	}

	if l.Insert(4, 2) {
		t.Error("Inserting at index out of boundaries should return false")
	}

	l.Insert(5, 1)
	l.Insert(6, 2)
	l.Insert(3, 1)
	l.Insert(8, 0)

	l2.Append(3)
	l2.Append(5)
	l2.Append(6)
	l2.Prepend(8)

	if !l.Compare(l2) {
		t.Error("Inserting failed")
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

func TestReverse(t *testing.T) {
	l := MakeList()
	l.Prepend(4)
	l.Prepend(5)
	l.Prepend(7)
	l.Prepend(2)
	l.Reverse()

	l2 := MakeList()
	l2.Append(4)
	l2.Append(5)
	l2.Append(7)
	l2.Append(2)

	if !l.Compare(l2) {
		t.Error("Reverse should reverse order of the list")
	}
}

func TestCompare(t *testing.T) {
	l := MakeList()
	l2 := MakeList()
	if !l.Compare(l2) {
		t.Error("Comparing empty lists failed")
	}
	l.Append(4)
	if l.Compare(l2) {
		t.Error("Comparing <4> and <nil> failed")
	}
	l2.Append(4)
	if !l.Compare(l2) {
		t.Error("Comparing <4> <4> failed")
	}
	l.Append(5)
	if l.Compare(l2) {
		t.Error("Comparing <4,5> <4> failed")
	}
	l2.Append(5)
	if !l.Compare(l2) {
		t.Error("Comparing <4,5> <4,5> failed")
	}
	l2.Append(7)
	l.Append(8)
	if l.Compare(l2) {
		t.Error("Comparing <4,5,7> <4,5,8> failed")
	}
}

func TestInsertAscending(t *testing.T) {
	l := MakeList()

	l.InsertAscending(4)
	l.InsertAscending(7)
	l.InsertAscending(3)
	l.InsertAscending(8)
	l.InsertAscending(0)
	l.InsertAscending(5)

	l2 := MakeList()
	l2.Append(0)
	l2.Append(3)
	l2.Append(4)
	l2.Append(5)
	l2.Append(7)
	l2.Append(8)

	if l.String() != l2.String() {
		t.Error("InsertAscending should insert items in the ascending order")
	}
}

func TestRemoveDuplicatesFromAscendingList(t *testing.T) {
	l := MakeList()
	l2 := MakeList()

	l.RemoveDuplicatesFromAscendingList()

	if !l.Compare(l2) {
		t.Error("RemoveDuplicates should not take actions on empty list")
	}
	l.InsertAscending(3)
	l.InsertAscending(3)
	l.InsertAscending(7)
	l.InsertAscending(3)
	l.RemoveDuplicatesFromAscendingList()

	l2.Prepend(3)
	l2.Append(7)

	if l.Compare(l2) {
		t.Error("RemoveDuplicates failure")
	}

}

func TestIsAscending(t *testing.T) {
	l := MakeList()
	if !l.IsAscending() {
		t.Error("IsAscending on empty list failure")
	}
	l.Prepend(2)
	if !l.IsAscending() {
		t.Error("IsAscending on list with one element failure")
	}
	l.Prepend(4)
	if l.IsAscending() {
		t.Error("IsAscending on descending list failure")
	}
	l.Prepend(1)
	if l.IsAscending() {
		t.Error("IsAscending on unsorted list failure")
	}
	l.Delete(4)
	if !l.IsAscending() {
		t.Error("IsAscending on ascending list failure #1")
	}
	l.Append(5)
	if !l.IsAscending() {
		t.Error("IsAscending on ascending list failure #2")
	}
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

func TestSortAscending(t *testing.T) {
	l := MakeList()
	l.Prepend(4)
	l.Prepend(7)
	l.Prepend(3)
	l.Prepend(8)
	l.Prepend(0)
	l.Prepend(5)

	l.SortAscending()

	l2 := MakeList()
	l2.Prepend(0)
	l2.Append(3)
	l2.Append(4)
	l2.Append(5)
	l2.Append(7)
	l2.Append(8)

	if !l.Compare(l2) {
		t.Error("Sorting failed!")
	}
}

func TestString(t *testing.T) {
	l := MakeList()
	if l.String() != "Empty list!" {
		t.Error("String on empty list failed")
		t.Error(l)
	}

	l.Prepend(5)
	l.Append(7)
	l.InsertAscending(6)
	l.Insert(4, 0)

	if l.String() != "->4->5->6->7->" {
		t.Error("String failure")
		t.Error(l)
	}
}

func TestFindMiddle(t *testing.T) {
	l := MakeList()
	if l.head.findMiddle() != nil {
		t.Error("Find middle failure on empty list")
	}

	l.Append(4)
	if l.head.findMiddle() != l.head {
		t.Error("Find middle on list with one node should return list's head")
	}

	l.Append(5)
	if l.head.findMiddle() != l.head {
		t.Error("Find middle on list with two nodes should return list's head")
	}

	l.Append(9)
	if l.head.findMiddle() != l.head.next {
		t.Error("Find middle on list with three nodes should return the second one")
	}

	l.Append(54)
	if l.head.findMiddle() != l.head.next {
		t.Error("Find middle on list with four nodes should return the second one")
	}

	l.Append(65)
	if l.head.findMiddle() != l.head.next.next {
		t.Error("Find middle on list with five nodes should return the third one")
	}
}

func TestMergeAscendingLists(t *testing.T) {
	n := makeNode(1)
	n2 := makeNode(2)
	if n.mergeAscendingLists(nil) != n {
		t.Error("merge list with nil failure")
	}

	n.next = makeNode(5)
	n2.next = makeNode(6)

	l := MakeList()
	l.head = n.mergeAscendingLists(n2)

	if l.String() != "->1->2->5->6->" {
		t.Error("Merge lists failure")
	}
}

func TestSortAscending2(t *testing.T) {
	var n *Node
	if n != n.sortAscending() {
		t.Error("sortAscending should return nil on empty list")
	}

	n = makeNode(4)
	if n != n.sortAscending() {
		t.Error("sortAscending on list with one node should return the node")
	}

	n.next = makeNode(6)
	n.next.next = makeNode(2)
	n.next.next.next = makeNode(1)

	n = n.sortAscending()

	l := MakeList()
	l.head = n

	if l.String() != "->1->2->4->6->" {
		t.Error("sortAscending should sort linked list in ascending order")
		t.Log(l)
	}
}
