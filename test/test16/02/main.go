package main

//
func searchMatrix(matrix [][]int, target int) bool {
	preI, preJ := -1, -1
	for i, j := len(matrix)-1, 0; i >= 0 && j < len(matrix[i]); {
		if i == preI && j == preJ {
			break
		}
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			j++
		} else {
			i--
		}
		preI, preJ = i, j
	}
	return false
}
