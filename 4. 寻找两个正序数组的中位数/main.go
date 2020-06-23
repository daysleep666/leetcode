package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var value, last int
	l := len(nums1) + len(nums2)
	odd := l%2 == 1
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		last = value
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				value = nums1[i]
				i++
			} else {
				value = nums2[j]
				j++
			}
		} else {
			if i < len(nums1) {
				value = nums1[i]
				i++
			}
			if j < len(nums2) {
				value = nums2[j]
				j++
			}
		}

		if odd { // 奇数
			if i+j-1 == l/2 {
				return float64(value)
			}
		} else { // 偶数
			if i+j-1 == l/2 {
				return (float64(value) + float64(last)) / 2
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5, 6}))
}
