package main

import "math"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	// 动态规划
	// 状态定义: d[i][j] 在(i,j)时的最短路径
	// 状态转移方程: d[i][j] = min(d[i-1][j],d[i][j-1])
	d := make([][]int, len(grid))
	for i := 0; i < len(d); i++ {
		d[i] = make([]int, len(grid[i]))
		for j := 0; j < len(d[i]); j++ {
			d[i][j] = int(math.MaxInt32)
		}
	}
	get := func(i, j int) int {
		if i < 0 || j < 0 {
			return int(math.MaxInt32)
		}
		return d[i][j]
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	d[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			d[i][j] = min(get(i-1, j), get(i, j-1)) + grid[i][j]
		}
	}
	return d[len(grid)-1][len(grid[0])-1]
}

func main() {

}
