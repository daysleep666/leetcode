package main

import "sort"

// 递归 + 剪枝
func largestPerimeter1(A []int) int {
	arr := make([]int, 0, 3)
	return recursive(A, 0, arr)
}

// 返回最大周长
func recursive(A []int, i int, arr []int) int {
	if len(arr) == 3 {
		if !valid(arr) {
			return 0
		}
		return arr[0] + arr[1] + arr[2]
	}
	if i == len(A) {
		return 0
	}

	skipArr := copyArr(arr)

	arr = append(arr, A[i])

	return max(recursive(A, i+1, arr), recursive(A, i+1, skipArr))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func copyArr(arr []int) []int {
	l := make([]int, len(arr))
	copy(l, arr)
	return l
}

func valid(arr []int) bool {
	return ((arr[0] + arr[1]) > arr[2]) && abs(arr[0]-arr[1]) < arr[2]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 排序加贪心
func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1; i >= 3; i++ {
		if valid(A[i-2:]) {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}

func main() {

}

// 给定由一些正数（代表长度）组成的数组 A，返回由其中三个长度组成的、面积不为零的三角形的最大周长。

// 如果不能形成任何面积不为零的三角形，返回 0。

// 示例 1：

// 输入：[2,1,2]
// 输出：5
// 示例 2：

// 输入：[1,2,1]
// 输出：0
// 示例 3：

// 输入：[3,2,3,4]
// 输出：10
// 示例 4：

// 输入：[3,6,2,3]
// 输出：8
