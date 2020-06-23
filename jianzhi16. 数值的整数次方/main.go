package main

import "fmt"

func myPow(x float64, n int) float64 {
	switch n {
	case 0:
		return 1
	case 1:
		return x
	case -1:
		return 1 / x
	}
	a := myPow(x, n/2)
	if n%2 == 0 {
		return a * a
	}
	if n < 1 {
		return a * a * (1 / x)
	}
	return a * a * x
}

func main() {
	fmt.Println(myPow(2, -2))
	fmt.Println(myPow(2, 5))
	fmt.Println(myPow(0.00001, 2147483647))
	fmt.Println(myPow(34.00515, -3))

}
