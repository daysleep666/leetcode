package main

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	v, c := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == v {
			c++
			continue
		}
		c--
		if c == 0 {
			v = nums[i]
			c = 1
		}
	}
	return v
}

func main() {

}
