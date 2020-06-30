package main

// 快排的partition
func removeElement(nums []int, val int) int {
	i, j := 0, 0
	k := 0
	for ; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
			k++
		}
	}
	return k
}

func main() {

}
