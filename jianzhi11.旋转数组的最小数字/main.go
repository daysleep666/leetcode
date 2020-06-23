package main

import "fmt"

func minArray1(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] <= numbers[i] {
			continue
		}
		return numbers[i]
	}
	return numbers[0]
}

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	low, high := 0, len(numbers)-1
	for low < high {
		mid := (low + high) / 2
		if numbers[mid] > numbers[high] {
			low = mid + 1
		} else if numbers[mid] < numbers[high] {
			high = mid
		} else {
			high--
		}

	}
	return numbers[low]
}
func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
	fmt.Println(minArray([]int{1, 3, 5}))
	fmt.Println(minArray([]int{3, 5, 1}))
	fmt.Println(minArray([]int{1, 1, 3, 1}))
	fmt.Println(minArray([]int{3, 1, 3}))
}
