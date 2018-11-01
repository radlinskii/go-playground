package singlylinked

import "testing"

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

func TestInsert(t *testing.T) {
	l := MakeList()
	l.Insert(3)

	if l.head.key != 3 {
		t.Error("Insert should insert 3 as key of list head")
	}

	l.Insert(5)
	if l.head.key != 5 {
		t.Error("Insert should insert 5 as key of list head")
	}

	if l.head.next.key != 3 {
		t.Error("Insert add new node at the begging of the list")
	}
}

func TestSearch(t *testing.T) {
	l := MakeList()

	if l.Search(3) != nil {
		t.Errorf("Search on empty list should definitely return nil")
	}
}
