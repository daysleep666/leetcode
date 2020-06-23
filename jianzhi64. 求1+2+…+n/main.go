package main

import "fmt"

func sumNums(n int) int {
	return (1 + n) * n / 2
}

func main() {
	fmt.Println(sumNums(3))
	fmt.Println(sumNums(9))
}
