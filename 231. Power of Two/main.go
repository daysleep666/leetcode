package main

func isPowerOfTwo1(n int) bool {
	if n == 0 {
		return false
	}
	for n != 1 {
		if n%2 != 0 {
			return false
		}
		n /= 2
	}
	return true
}

func isPowerOfTwo(n int) bool {
	return n != 0 && (n&(n-1)) == 0
}
