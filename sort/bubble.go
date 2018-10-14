// Package sort contains various implementations of sorting algorithms.
package sort

// Bubble is an implementation of notorious bubble sort.
func Bubble(s []int) []int {
	n := len(s)
	for ; n > 0; n-- {
		for i := 0; i < n-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
	return s
}
