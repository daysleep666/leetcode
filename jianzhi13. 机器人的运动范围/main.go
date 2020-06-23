package main

import "fmt"

// dfs
// 机器人走的路线可能不是一个对称的图案 9+9 / 10+10
func movingCount(m int, n int, k int) int {
	count := 0
	board := make([][]int, m)
	for i := 0; i < m; i++ {
		board[i] = make([]int, n)
	}
	recursive(board, m, n, k, 0, 0, &count)
	return count
}

func recursive(board [][]int, m, n, k, i, j int, count *int) {
	if i >= m || j >= n {
		return
	}
	// 已经路过这个点了
	if board[i][j] == '/' {
		return
	}
	// 确认下当前点是否可达
	if toSum(i)+toSum(j) > k {
		return
	}
	board[i][j] = '/'
	(*count)++
	// 只会往下或往右
	recursive(board, m, n, k, i+1, j, count)
	recursive(board, m, n, k, i, j+1, count)
}

func toSum(v int) int {
	sum := 0
	for v != 0 {
		sum += v % 10
		v /= 10
	}

	return sum
}

func main() {
	fmt.Println(movingCount(30, 30, 3))
}
