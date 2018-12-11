package ctci

// ZeroMatrix modifies a matrix so that if it finds a zero in the matrix,
// it turns its whole row and column into zeros. Works for Matrixes of dimension MxN.
func ZeroMatrix(arr [][]int) {
	m := len(arr)
	if m > 0 {
		n := len(arr[0])

		rows := make([]bool, m)
		cols := make([]bool, n)

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if arr[i][j] == 0 {
					rows[i] = true
					cols[j] = true
				}
			}
		}
		// "zero out" rows
		for i := 0; i < m; i++ {
			if rows[i] == true {
				for j := 0; j < n; j++ {
					arr[i][j] = 0
				}
			}
		}
		// "zero out" columns
		for i := 0; i < n; i++ {
			if cols[i] == true {
				for j := 0; j < m; j++ {
					arr[j][i] = 0
				}
			}
		}
	}
}
