package main

import "fmt"

func maxProduct(nums []int) int {
	return recusion(nums[1:], map[int]int{}, nums[0])
}

func recusion(n []int, m map[int]int, v int) int {
	if len(n) == 0 {
		return v
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	get := func(c int) int {
		if a, ok := m[c]; ok {
			return a
		}
		a := recusion(n[1:], m, c)
		m[c] = a
		return a
	}

	a := get(n[0])
	b := get(n[0] * v)
	return max(max(a, b), v)
}

func maxProduct2(nums []int) int {
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	getMin := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	result, max, min := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if (min*nums[i]) > (max*nums[i]) && (min*nums[i]) > (nums[i]) {
			min, max = getMin(max*nums[i], nums[i]), min*nums[i]
		} else {
			max, min = getMax(max*nums[i], nums[i]), getMin(getMin(min*nums[i], nums[i]), max*nums[i])
		}
		result = getMax(result, max)
	}
	return result
}

func main() {
	// fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	// fmt.Println(maxProduct([]int{-2, 0, -1}))
	// fmt.Println(maxProduct([]int{-2, 3, -4}))
	// fmt.Println(maxProduct([]int{2, -5, -2, -4, 3}))
	// fmt.Println(maxProduct([]int{2, 3, -1, 1, -3, 3, 0}))
	fmt.Println(maxProduct([]int{1, 2, -1, -2, 2, 1, -2, 1}))
}
