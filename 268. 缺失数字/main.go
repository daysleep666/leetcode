package main

func missingNumber(nums []int) int {
	var bitmask int
	for i, v := range nums {
		bitmask ^= i
		bitmask ^= v
	}
	bitmask ^= len(nums)
	return bitmask
}

func main() {

}
