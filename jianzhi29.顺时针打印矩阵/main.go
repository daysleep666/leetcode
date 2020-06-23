package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	count := len(matrix) * len(matrix[0])
	result := make([]int, 0, count)

	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	k := 0

	add := func(i, j int) {
		result = append(result, matrix[i][j])
		k++
	}
	for left < right && top < bottom {
		for i := left; i < right; i++ {
			add(top, i)
		}
		for i := top; i < bottom; i++ {
			add(i, right)
		}
		for i := right; i > left; i-- {
			add(bottom, i)
		}

		for i := bottom; i > top; i-- {
			add(i, left)
		}
		top++
		bottom--
		left++
		right--
	}
	if top == bottom {
		for i := left; i <= right; i++ {
			add(top, i)
		}
	} else if left == right {
		for i := top; i <= bottom; i++ {
			add(left, i)
		}
	}

	return result
}

func main() {
	// fmt.Println(spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
	// fmt.Println(spiralOrder([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}))
	fmt.Println(spiralOrder([][]int{[]int{3, 2}}))
}

// 1 2 3
// 4 5 6
// 7 8 9
