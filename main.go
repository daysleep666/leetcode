package main

func main() {
	add := func(v int) {
		if len(arr) == 0 {
			arr = append(arr, v)
			return
		}
		low, high, mid := 0, len(arr)-1, 0
		for low <= high {
			// 超出int64
			mid = (low + high) / 2
			if arr[mid] == v {
				low = mid
				break
			} else if arr[mid] > v {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		arr = append(arr, 0)
		for i := len(arr) - 1; i > low; i-- {
			arr[i] = arr[i-1]
		}
		arr[low] = v
	}
}
