package main

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	visited := make(map[string]bool, 0)
	recursive(nums, 0, visited, &result)

	return result
}

func recursive(nums []int, i int, visited map[string]bool, result *[][]int) {
	for j := i; j < len(nums); j++ {
		arr := copyArr(nums)
		arr[i], arr[j] = arr[j], arr[i]
		key := ints2String(arr)
		if !visited[key] {
			visited[key] = true
			(*result) = append((*result), arr)
		}
		recursive(arr, i+1, visited, result)
	}
}

func copyArr(nums []int) []int {
	arr := make([]int, len(nums))
	copy(arr, nums)
	return arr
}

func ints2String(nums []int) string {
	str := make([]byte, len(nums))
	for i, v := range nums {
		str[i] = byte(v)
	}
	return string(str)
}
