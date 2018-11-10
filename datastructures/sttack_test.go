package datastructures

import "testing"

func TestMakeStack(t *testing.T) {
	s := MakeStack()
	if !s.list.IsEmpty() {
		t.Error("MakeStack should create a new stack with empty list")
	}
}

func TestPush(t *testing.T) {
	s := MakeStack()
	s.Push(4)
	if s.list.GetHead().GetKey() != 4 {
		t.Error("Push should add element to the empty stack")
	}
	s.Push(5)
	if s.list.GetHead().GetKey() != 5 {
		t.Error("Push should add element with given value at the begginning of the stack")
	}
}

func TestPop(t *testing.T) {
	s := MakeStack()
	s.Push(4)
	s.Push(5)

	n := s.Pop()
	if n.GetKey() != 5 {
		t.Error("Pop should return the last value that was inserted to the stack")
	}
	n = s.Pop()
	if n.GetKey() != 4 {
		t.Error("Pop should remove element from the stack")
	}
	n = s.Pop()
	if n != nil {
		t.Error("Pop on empty stack should return nil")
	}
}

func TestIsStackEmpty(t *testing.T) {
	s := MakeStack()
	if !s.IsEmpty() {
		t.Error("IsEmpty should return true on empty stack")
	}

	s.Push(5)
	if s.IsEmpty() {
		t.Error("IsEmpty should return false when there are elements in the stack")
	}
}

func TestTop(t *testing.T) {
	s := MakeStack()
	if s.Top() != nil {
		t.Error("Top on empty stack should return nil")
	}

	s.Push(4)
	s.Push(5)
	if s.Top().GetKey() != 5 {
		t.Error("Top should return first node from the stack")
	}

	if s.Top().GetKey() != 5 {
		t.Error("Top should not remove nodes from the stack")
	}
}

func TestStackString(t *testing.T) {
	s := MakeStack()
	got := s.String()
	expected := "Empty stack!"
	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}

	s.Push(4)
	s.Push(5)
	s.Push(7)
	got = s.String()
	expected = "->7.00->5.00->4.00->"
	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}

	s.Pop()
	s.Pop()

	got = s.String()
	expected = "->4.00->"
	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}

}
