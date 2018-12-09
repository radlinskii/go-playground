package ctci

// LeftRotate rotates the matrix by 90 degrees.
func LeftRotate(arr [][]int) {
	if len(arr) > 1 {
		for i := 0; i <= len(arr)/2; i++ {
			for j := i; j+i < len(arr)-1; j++ {
				tmp1 := arr[i][j]
				arr[i][j] = arr[len(arr)-1-j][i]
				tmp2 := arr[j][len(arr)-1-i]
				arr[j][len(arr)-1-i] = tmp1
				tmp1 = arr[len(arr)-1-i][len(arr)-1-j]
				arr[len(arr)-1-i][len(arr)-1-j] = tmp2
				arr[len(arr)-1-j][i] = tmp1
			}
		}
	}
}
