package main

import "fmt"

func smallestK(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	if len(arr) == k {
		return arr
	}

	i := 0
	for j := 0; j < len(arr)-1; {
		if arr[j] > arr[len(arr)-1] {
			j++
			continue
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j++
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]

	if len(arr[:i]) == k {
		return arr[:i]
	} else if len(arr[:i]) < k {
		return append(arr[:i], smallestK(arr[i:], k-len(arr[:i]))...)
	}
	return smallestK(arr[:i], k)
}

func main() {
	arr := []int{9, 1, 3, 5, 7, 2, 4, 6, 8}
	fmt.Println(smallestK(arr, 4))
}
