package main

func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp := []int{triangle[0][0]}
	get := func(i int) int {
		if i < 0 {
			return dp[0]
		} else if i >= len(dp) {
			return dp[len(dp)-1]
		}
		return dp[i]
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 1; i < len(triangle); i++ {
		tmp := make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			tmp[j] = triangle[i][j] + min(get(j-1), get(j))
		}
		dp = tmp
	}
	m := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] < m {
			m = dp[i]
		}
	}
	return m
}

func minimumTotal(triangle [][]int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}

func main() {

}
