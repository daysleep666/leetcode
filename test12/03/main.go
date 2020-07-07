package main

func rotate(matrix [][]int) {
	// 1. 镜像
	for i := 0; i < len(matrix); i++ {
		reverse(matrix[i])
	}
	// 2.以斜线镜像
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
		}
	}
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
