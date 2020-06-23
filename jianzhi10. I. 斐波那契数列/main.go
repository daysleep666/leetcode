package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	d := make([]int, n+1)
	d[0] = 0
	d[1] = 1
	for i := 2; i <= n; i++ {
		d[i] = (d[i-1] + d[i-2]) % (1e9 + 7)
	}
	return d[n]
}

func main() {
	fmt.Println(fib(45))
	fmt.Println(fib(95))
}
