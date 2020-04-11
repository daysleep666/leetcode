package main

import "fmt"

func main() {
	list := []int{1, 4, 6, 9, 12}
	getMid := func(v int) int {
		low, high := 0, len(list)
		mid := 0
		for low < high {
			mid = (low + high) / 2
			if list[mid] == v {
				return mid
			} else if list[mid] < v {
				low = mid + 1
			} else {
				high = mid
			}
		}
		return low
	}
	fmt.Println(getMid(2))
	fmt.Println(getMid(3))
	fmt.Println(getMid(5))
	fmt.Println(getMid(7))
	fmt.Println(getMid(8))
	fmt.Println(getMid(10))
	fmt.Println(getMid(11))
	fmt.Println(getMid(13))
	fmt.Println(getMid(0))

}
