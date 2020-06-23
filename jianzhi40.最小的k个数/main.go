package main

import "fmt"

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 {
		return nil
	}
	if k == 0 {
		return nil
	}
	if k == len(arr) {
		return arr
	}
	list := make([]int, k)
	copy(list, arr)
	heapify := func(i int) {
		cur := i
		for cur < len(list) {
			if i*2+1 < len(list) && list[cur] < list[i*2+1] {
				cur = i*2 + 1
			}
			if i*2+2 < len(list) && list[cur] < list[i*2+2] {
				cur = i*2 + 2
			}
			if i == cur {
				return
			}
			list[i], list[cur] = list[cur], list[i]
			i = cur
		}
	}

	for i := len(list) / 2; i >= 0; i-- {
		heapify(i)
	}
	for i := k; i < len(arr); i++ {
		if arr[i] > list[0] {
			continue
		}
		list[0] = arr[i]
		heapify(0)
	}
	return list
}

func main() {
	fmt.Println(getLeastNumbers([]int{0, 1, 1, 2, 4, 4, 1, 3, 3, 2}, 6))
}
