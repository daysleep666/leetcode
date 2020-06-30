package main

func permute(nums []int) [][]int {
	result := make([][]int, 0, len(nums)*len(nums))
	recursive(nums, 0, &result)
	return result
}

func recursive(nums []int, first int, result *[][]int) {
	// 终止条件
	if first == len(nums) {
		(*result) = append((*result), nums)
		return
	}
	//
	for i := first; i < len(nums); i++ {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		tmp[first], tmp[i] = tmp[i], tmp[first]
		recursive(tmp, first+1, result)
	}
}

func main() {

}
