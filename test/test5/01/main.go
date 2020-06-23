package main

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	// 动态规划
	// 状态定义: d[n] 到第n阶台阶时的走法数量
	// 状态转移方程: d[n] = d[n-1] + d[n-2]
	// 初始状态
	d := make([]int, n+1)
	d[0] = 1
	d[1] = 1
	for i := 2; i <= n; i++ {
		d[i] = d[i-1] + d[i-2]
	}
	return d[n]
}

func main() {

}
