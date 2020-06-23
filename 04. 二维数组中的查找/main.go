package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i, j := len(matrix)-1, 0; i >= 0 && j < len(matrix[i]); {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return false
}

func main() {

}
