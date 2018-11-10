package algo

import (
	"github.com/radlinskii/go-playground/datastructures"
)

// IsBalanced takes a string containing brackets
// and checks if the brackets are in balanced order.
func IsBalanced(s string) bool {
	stack := datastructures.MakeStack()
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	stack.Push(float32(s[0]))
	for _, v := range s[1:len(s)] {
		if v == 123 || v == 40 || v == 91 {
			stack.Push(float32(v))
		} else {
			if v == 125 { // }
				if stack.IsEmpty() || stack.Top().GetKey() != 123 { // {
					return false
				}
				stack.Pop()
			} else if v == 41 { // )
				if stack.IsEmpty() || stack.Top().GetKey() != 40 { // (
					return false
				}
				stack.Pop()
			} else if v == 93 { // ]
				if stack.IsEmpty() || stack.Top().GetKey() != 91 { // [
					return false
				}
				stack.Pop()
			}
		}
	}
	return stack.IsEmpty()
}
