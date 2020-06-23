package main

import "fmt"

func lastRemaining(n int, m int) int {
	alive := 0
	for i := 2; i <= n; i++ {
		alive = (alive + m) % i
	}
	return alive
}

func main() {
	fmt.Println(lastRemaining(5, 3))
	fmt.Println(lastRemaining(100000, 100))
}

// m = 3
// 1  -- 0
// 1,2 -- 2
// 1,2,3 -- 1,2
// 1,2,3,4 -- 1,2,4
