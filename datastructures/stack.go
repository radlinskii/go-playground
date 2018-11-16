package datastructures

import (
	"github.com/radlinskii/go-playground/datastructures/lists/singlylinked"
)

// Stack type is a implementation of stack data structure.
// It uses singly linked list datastructure and its functionalities.
type Stack struct {
	list *singlylinked.List
}

// MakeStack creates an empty stack.
func MakeStack() *Stack {
	return &Stack{list: singlylinked.MakeList()}
}

// Push inserts an element of given value at the top of the stack.
func (s *Stack) Push(v float32) {
	s.list.Prepend(v)
}

// Pop returns an element from the top after removing it from the stack.
func (s *Stack) Pop() *singlylinked.Node {
	return s.list.Pop()
}

// IsEmpty checks if top of the stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.list.IsEmpty()
}

// Top returns first element from the stack witthout removing it.
func (s *Stack) Top() *singlylinked.Node {
	return s.list.GetHead()
}

func (s *Stack) String() string {
	if s.list.IsEmpty() {
		return "Empty stack!"
	}
	return s.list.String()
}

// SortAscending sorts the elements of the stack in ascending order.
func (s *Stack) SortAscending() {
	tmp := MakeStack()
	var curr *singlylinked.Node
	for !s.IsEmpty() {
		curr = s.Pop()
		for !tmp.IsEmpty() && tmp.Top().GetKey() <= curr.GetKey() {
			s.Push(tmp.Pop().GetKey())
		}
		tmp.Push(curr.GetKey())

	}
	s.list = tmp.list
}
