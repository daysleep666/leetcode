package pkg

import "sort"

// BinarySearch 二分查找
func BinarySearch(list []int, v int) bool {
	sort.Ints(list)
	for low, high := 0, len(list)-1; low <= high; {
		mid := (low + high) / 2
		midValue := list[mid]
		if v == midValue {
			return true
		} else if v > midValue {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
