package main

import "fmt"

func isBadVersion(version int) bool {
	if version >= 1 {
		return true
	}
	return false
}

// 二分
func firstBadVersion(n int) int {
	start, end := 1, n
	for start < end {
		mid := (start + end) / 2
		bad := isBadVersion(mid)
		// fmt.Println(start, end, mid, bad)
		if bad {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

/*



 */

func main() {
	fmt.Println(firstBadVersion(5))
}
