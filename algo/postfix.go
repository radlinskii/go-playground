package algo

import (
	"strconv"
	"strings"

	"github.com/radlinskii/go-playground/datastructures"
)

// ReversePolishNotation reads and evaluates given expression.
// The provided string should be a valid RPN expression.
func ReversePolishNotation(exprStr string) float32 {
	s := datastructures.MakeStack()

	expr := strings.Split(exprStr, " ")

	for _, c := range expr {
		if num, err := strconv.ParseFloat(c, 32); err == nil {
			s.Push(float32(num))
		} else {
			v1 := s.Pop()
			v2 := s.Pop()

			switch int(c[0]) {
			case 43: // +
				s.Push(v2.GetKey() + v1.GetKey())
			case 45: // -
				s.Push(v2.GetKey() - v1.GetKey())
			case 42: // *
				s.Push(v2.GetKey() * v1.GetKey())
			case 47: // /
				s.Push(v2.GetKey() / v1.GetKey())
			}
		}
	}
	return s.Pop().GetKey()
}
