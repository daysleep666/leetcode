package main

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	// [1,len(nums)] 答案一定在这个区间中 所以将负数和大于这个数的数据都置0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 1 || nums[i] > len(nums) {
			nums[i] = 0
		}
	}

	// 所有数都在[1,len(nums)]中

	// 将数组作为哈希表
	for i := 0; i < len(nums); i++ {
		for {
			if nums[i] == i+1 || nums[i] == 0 {
				break
			}
			index := nums[i] - 1
			if nums[i] == nums[index] {
				break
			}
			nums[i], nums[index] = nums[index], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums)
}

func main() {

}
