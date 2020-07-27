package main

import "math"

func shortestSubarray(A []int, K int) int {
	minLen := int(math.MaxInt32)
	for i := 0; i < len(A); i++ {
		sum := 0
		for j := i; j < len(A); j++ {
			if sum += A[j]; sum == K {
				minLen = min(minLen, j-i+1)
				break
			}
		}
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
