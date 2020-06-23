package pkg

// BinarySearch 二分查找
func BinarySearch(arr []int, v int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + high
		if arr[mid] == v {
			return mid
		} else if arr[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// BinarySearch1 变种二分查找: 大于某值且最小
func BinarySearch1(arr []int, v int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + high
		if arr[mid] <= v {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

// BinarySearch2 变种二分查找: 小于某值且最大
func BinarySearch2(arr []int, v int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + high
		if arr[mid] < v {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}
