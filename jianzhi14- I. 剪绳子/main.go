package main

import "fmt"

func cuttingRope(n int) int {
	if n == 2 {
		return 1
	}
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max := 0
	for i := 2; i < n; i++ {
		max = getMax(cut(n, i, 1), getMax(cut(n, i, 0), max))
	}
	return max
}

func cut(n, m int, add int) int {
	mul := 1
	sum := 0
	for i := 0; i < m-1; i++ {
		a := n/m + add
		sum += a
		mul *= a
	}
	mul *= n - sum
	return mul
}

func main() {
	// fmt.Println("--->", cut(8, 3))
	fmt.Println(cuttingRope(3))  // 2
	fmt.Println(cuttingRope(8))  // 18
	fmt.Println(cuttingRope(10)) // 36
}
