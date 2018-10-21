package sort

// Insert is an implementation of humble insert sort.
func Insert(s []int) []int {
	n := len(s)
	for i := 1; i < n; i++ {
		x := s[i]
		j := i - 1
		for j >= 0 && s[j] > x {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = x
	}
	return s
}
