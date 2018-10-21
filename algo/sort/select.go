package sort

// Select is an implementation of simple select sort algorithm.
func Select(s []int) []int {
	n := len(s)
	for i := 0; i < n; i++ {
		x := i
		for j := i + 1; j < n; j++ {
			if s[x] > s[j] {
				x = j
			}
		}
		s[x], s[i] = s[i], s[x]
	}
	return s
}
