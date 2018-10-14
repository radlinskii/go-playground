package sort

// Merge is an implementation of divide-and-conquare merge sort algorithm.
func Merge(s []int) []int {
	n := len(s)
	if n >= 2 {
		q := n / 2
		s1 := Merge(s[:q])
		s2 := Merge(s[q:])
		return merge(s1, s2)
	}
	return s
}

func merge(s1, s2 []int) []int {
	len1, len2 := len(s1), len(s2)
	s3 := make([]int, len1+len2)
	i, j := 0, 0

	for i < len1 && j < len2 {
		if s1[i] < s2[j] {
			s3[i+j] = s1[i]
			i++
		} else {
			s3[i+j] = s2[j]
			j++
		}
	}

	for i < len1 {
		s3[i+j] = s1[i]
		i++
	}

	for j < len2 {
		s3[i+j] = s2[j]
		j++
	}

	return s3
}
