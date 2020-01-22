package main

func threeSum(nums []int) (result [][]int) {
	sort := func(l []int) { //  min mid max
		if l[0] > l[1] {
			l[0], l[1] = l[1], l[0]
		}
		if l[2] < l[0] {
			l[0], l[1], l[2] = l[2], l[0], l[1]
		} else if l[2] < l[1] {
			l[1], l[2] = l[2], l[1]
		}
	}

	add := func(l []int) {
		sort(l)
		for _, v := range result {
			if v[0] == l[0] && v[1] == l[1] && v[2] == l[2] {
				return
			}
		}
		result = append(result, l)
	}

	return
}

func main() {
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

// Note:

// The solution set must not contain duplicate triplets.

// Example:

// Given array nums = [-1, 0, 1, 2, -1, -4],

// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
