package main

func getPermutation(n int, k int) string {
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}
	recursive(arr, k, 0)

	str := make([]byte, n)
	for i, v := range arr {
		str[i] = byte('0' + v)
	}
	return string(str)
}

func recursive(arr []int, k, first int) {
	if len(arr) == first {
		return
	}
	base := factorial(len(arr) - first - 1)
	cur := 0
	for i := first; i < len(arr); i++ {
		arr[first], arr[i] = arr[i], arr[first]
		if k <= cur+base {
			recursive(arr, k-cur, first+1)
			return
		}

		cur += base
	}
}

func factorial(n int) int {
	for i := n - 1; i >= 1; i-- {
		n *= i
	}
	return n
}

func main() {

}

// 1 2 3
// 1 3 2
// 2 1 3
// 2 3 1
// 3 1 2
// 3 2 1
