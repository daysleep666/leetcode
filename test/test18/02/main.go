package main

func integerBreak(n int) int {

	// 2-n-1段
	result := 0
	for i := 2; i < n; i++ {
		result = max(result, divide(n, 2))
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func divide(num int, c int) int {

}
