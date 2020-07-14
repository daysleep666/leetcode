package main

func sortedSquares(A []int) []int {
	quickSort(A, 0, len(A)-1)
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}
	return A
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := partition(arr, start, end)
	quickSort(arr, start, mid-1)
	quickSort(arr, mid+1, end)
}

func partition(arr []int, start, end int) int {
	i, j := start, start
	for ; i < end; i++ {
		if abs(arr[i]) < abs(arr[end]) {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	arr[end], arr[j] = arr[j], arr[end]
	return j
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
