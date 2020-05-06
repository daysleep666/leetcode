package main

import "fmt"

func isPossible(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	list := [][]int{[]int{nums[0]}}
	canAppend := func(arr []int, v int) bool {
		l := len(arr)
		if arr[l-1]+1 == v {
			return true
		}
		return false
	}
	add := func(v int) {
		i := len(list) - 1
		for ; i >= 0; i-- {
			if canAppend(list[i], v) {
				list[i] = append(list[i], v)
				return
			}
		}
		if i == -1 {
			list = append(list, []int{v})
		}
	}
	for i := 1; i < len(nums); i++ {
		add(nums[i])
	}
	for _, arr := range list {
		if len(arr) < 3 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPossible([]int{1, 2, 3}))
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 4, 5, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 4, 4, 5}))
}
