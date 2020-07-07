package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := m[v]; ok {
			return []int{i, j}
		}
		m[target-v] = i
	}
	return nil
}
