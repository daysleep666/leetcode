package main

import "fmt"

// 递归 + 剪枝
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			sum += recursive(matrix, i, j, 0, target, 0, 1)
			sum += recursive(matrix, i, j, 0, target, 1, 0)
			sum += recursive1(matrix, i, j, 0, target, 1, 1)
		}
	}
	return sum
}

func recursive(matrix [][]int, i, j int, sum, target int, a, b int) int {
	if i >= len(matrix) || j >= len(matrix[0]) {
		return 0
	}

	sum += matrix[i][j]
	if sum == target {
		fmt.Println(i, j)

		return 1
	}

	// 只能往右,只能往下,同时右下
	return recursive(matrix, i+a, j+b, sum, target, a, b)
}

func recursive1(matrix [][]int, i, j int, sum, target int, a, b int) int {
	if i >= len(matrix) || j >= len(matrix[0]) {
		return 0
	}

	sum = 0
	for m := 0; m <= i; m++ {
		for n := 0; n <= j; n++ {
			sum += matrix[m][n]
		}
	}

	if sum == target {
		return 1
	}

	// 只能往右,只能往下,同时右下
	return recursive(matrix, i+a, j+b, sum, target, a, b)
}

func main() {
	fmt.Println(numSubmatrixSumTarget([][]int{[]int{1, -1}, []int{-1, 1}}, 0))
}
