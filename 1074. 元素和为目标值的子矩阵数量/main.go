package main

import "fmt"

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	get := func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		return matrix[i][j]
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = get(i-1, j) + get(i, j-1) - get(i-1, j-1) + matrix[i][j]
		}
	}
	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for m := i; m < len(matrix); m++ {
				for n := j; n < len(matrix[i]); n++ {
					if v := matrix[m][n] - get(m, j-1) - get(i-1, n) + get(i-1, j-1); v == target {
						result++
					}
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(numSubmatrixSumTarget([][]int{
		[]int{0, 1, 0},
		[]int{1, 1, 1},
		[]int{0, 1, 0},
	}, 0))
}
