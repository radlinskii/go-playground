package algo

import (
	"github.com/radlinskii/go-playground/datastructures/lists/singlylinked"
)

// IsBalanced takes a string containing brackets
// and checks if the brackets are in balanced order.
func IsBalanced(s string) bool {
	list := singlylinked.MakeList()
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	list.Prepend(int(s[0]))
	for _, v := range s[1:len(s)] {
		if v == 123 || v == 40 || v == 91 {
			list.Prepend(int(v))
		} else {
			if v == 125 { // }
				if list.IsEmpty() || list.GetHead().GetKey() != 123 { // {
					return false
				}
				list.Pop()
			} else if v == 41 { // )
				if list.IsEmpty() || list.GetHead().GetKey() != 40 { // (
					return false
				}
				list.Pop()
			} else if v == 93 { // ]
				if list.IsEmpty() || list.GetHead().GetKey() != 91 { // [
					return false
				}
				list.Pop()
			}
		}
	}
	return list.IsEmpty()
}
