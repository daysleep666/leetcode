package main

func exchange(nums []int) []int {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

func main() {

}
