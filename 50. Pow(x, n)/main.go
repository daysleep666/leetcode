package main

import "fmt"

func myPow(x float64, n int) float64 {
	switch {
	case n == 0:
		return 1
	case n < 0:
		return 1 / myPow(x, -n)
	case n%2 != 0:
		return x * myPow(x*x, n/2)
	default:
		return myPow(x*x, n/2)
	}
}

func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	tmpN := n
	if n < 0 {
		n = -n
	}
	tmp := float64(1)
	for n != 1 {
		// n&1==1 奇数
		if n&1 == 1 {
			tmp *= x
		}
		x = x * x
		n = n / 2
	}
	x *= tmp
	if tmpN < 0 {
		x = 1 / x
	}
	return x
}

func main() {
	fmt.Println(myPow1(2, 10))
	fmt.Println(myPow1(2.1, 3))
	fmt.Println(myPow1(2, -2))
}
