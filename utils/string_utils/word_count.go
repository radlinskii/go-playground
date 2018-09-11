package string_utils

import "strings"

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	f := strings.Fields(s)

	for _, v := range f {
		_, ok := m[v]
		if ok == true {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	return m
}
