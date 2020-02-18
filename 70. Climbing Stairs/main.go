package main

func climbStairs1(n int) int {
	return d(n, map[int]int{})
}

func d(n int, r map[int]int) int {
	if n <= 2 {
		return n
	}

	f := func(n int) int {
		if v, ok := r[n]; ok {
			return v
		}
		v := d(n, r)
		r[n] = v
		return v
	}

	a := f(n - 1)
	b := f(n - 2)
	return a + b
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {

}
