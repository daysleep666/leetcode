package main

import "fmt"

func main() {
	// fmt.Println(getHappyString(1, 1))
	// fmt.Println(getHappyString(1, 3))
	// fmt.Println(getHappyString(2, 1))
	// fmt.Println(getHappyString(2, 3))
	// fmt.Println(getHappyString(2, 4))
	fmt.Println(getHappyString(3, 4))    // acb
	fmt.Println(getHappyString(4, 20))   // "cacb"
	fmt.Println(getHappyString(7, 100))  // "bcabacb"
	fmt.Println(getHappyString(10, 100)) // abacbabacb
}

func getHappyString(n int, k int) string {
	arr := make([]byte, 0, n)
	get := func(i int) byte {
		switch i {
		case 1:
			return 'a'
		case 2:
			return 'b'
		case 3:
			return 'c'
		}
		return ' '
	}
	add := func(i int) {
		if len(arr) == 0 {
			arr = append(arr, get(i))
			return
		}
		v := arr[len(arr)-1] - 'a' + 1
		if v == 1 {
			i += int(v)
		} else if v == 2 {
			if i == 2 {
				i = 3
			}
		}
		arr = append(arr, get(i))
	}
	for i := n; i > 0; i-- {
		r := pow(2, i-1)
		sum := r
		j := 1
		for ; k > sum; j++ {
			sum += r
		}
		// fmt.Println(j, k)
		if j > 3 {
			return ""
		}
		add(j)
		k = k - (sum - r)
	}
	return string(arr)
}

func pow(a, b int) int {
	mul := 1
	for i := 0; i < b; i++ {
		mul *= a
	}
	return mul
}

/*
n=1 a,b,c
n=2 ab,ac,ba,bc,ca,cb  3*2

3*2*2 = 12

cab

"aba"
"abc"
"aca"
"acb"
"bab"
"bac"
"bca"
"bcb"
"cab"
"cac"
"cba"
"cbc"



n=10
3*2*2*2*2*2*2*2*2*2

*/
