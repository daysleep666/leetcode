package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := nums[0]
	max := sum
	for i := 1; i < len(nums); i++ {
		if sum+nums[i] > nums[i] {
			sum = sum + nums[i]
		} else {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {

}
