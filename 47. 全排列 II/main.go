package main

import "sort"

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	recursive(nums, 0, &result)
	return result
}

func recursive(nums []int, index int, result *[][]int) {
	if index == len(nums)-1 {
		(*result) = append((*result), nums)
		return
	}
	for i := index; i < len(nums); i++ {
		if index < i && nums[index] == nums[i-1] {
			continue
		}
		arr := copyArr(nums)
		arr[i], arr[index] = arr[index], arr[i]
		recursive(arr, index+1, result)
	}
}

func permuteUnique1(nums []int) [][]int {
	result := make([][]int, 0)
	recursive1(nums, 0, &result)
	return result
}

func recursive1(nums []int, index int, result *[][]int) {
	if index == len(nums)-1 {
		(*result) = append((*result), nums)
		return
	}
	m := make(map[int]bool, 0)
	for i := index; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		arr := copyArr(nums)
		arr[i], arr[index] = arr[index], arr[i]
		m[nums[i]] = true
		recursive(arr, index+1, result)
	}
}

func copyArr(nums []int) []int {
	arr := make([]int, len(nums))
	copy(arr, nums)
	return arr
}

func main() {

}
