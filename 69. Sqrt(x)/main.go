package main

import "fmt"

func mySqrt(x int) int {
	low, high := 0, x
	for low <= high {
		mid := (low + high) / 2
		v := mid * mid
		if v == x {
			return mid
		} else if v > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(10))
}
