package algo

import (
	"github.com/radlinskii/go-playground/utils/str"
)

// Convert converts number given in string
// from the current system to the new one.
// 2 <= curr, new <= 16
func Convert(curr, new int, number string) string {
	var nums []int
	for _, c := range number {
		if c < 'A' {
			nums = append(nums, int(c-'0'))
		} else {
			nums = append(nums, 10+int(c-'A'))
		}
	}

	dec := toDecimal(curr, nums)

	res := fromDecimal(new, dec)

	return str.Reverse(res)
}

func toDecimal(curr int, nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = nums[i] + res*curr
	}
	return res
}

// returns string with number in new system but in reversed order.
func fromDecimal(new, dec int) string {
	res := ""
	if dec == 0 {
		return "0"
	}
	for dec != 0 {
		newC := dec % new
		if newC < 10 {
			res += string('0' + newC)
		} else {
			res += string('A' + newC - 10)
		}
		dec /= new
	}

	return res
}
