package main

func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))
	// 单调栈
	stack := make([]int, 0, len(T))
	for i := 0; i < len(T); i++ {
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if T[top] >= T[i] {
				break
			}
			result[top] = i - top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}
