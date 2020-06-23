package main

import "fmt"

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	d := make([]int, n+1)
	d[1] = 1
	d[2] = 2
	for i := 3; i <= n; i++ {
		d[i] = (d[i-1] + d[i-2]) % (1e9 + 7)
	}
	return d[n]
}

func main() {
	fmt.Println(numWays(7))
}
