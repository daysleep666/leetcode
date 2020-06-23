package main

func maxValue(grid [][]int) int {
	// 动态规划
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	get := func(i, j int) int {
		if i < 0 {
			return 0
		}
		if j < 0 {
			return 0
		}
		return grid[i][j]
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = max(get(i-1, j), get(i, j-1)) + grid[i][j]
		}
	}

	// 右下角一定是最大值
	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {

}
