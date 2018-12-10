package ctci

// LeftRotate rotates the matrix by 90 degrees.
func LeftRotate(arr [][]int) {
	if len(arr) > 1 {
		for i := 0; i <= len(arr)/2; i++ {
			for j := i; j+i < len(arr)-1; j++ {
				tmp := arr[i][j]
				arr[i][j] = arr[len(arr)-1-j][i]
				arr[len(arr)-1-j][i] = arr[len(arr)-1-i][len(arr)-1-j]
				arr[len(arr)-1-i][len(arr)-1-j] = arr[j][len(arr)-1-i]
				arr[j][len(arr)-1-i] = tmp
			}
		}
	}
}
