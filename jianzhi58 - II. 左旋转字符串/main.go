package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	//
	n = n % len(s)
	//
	list := []byte(s)
	tmp := list[:n]
	list = list[n:]
	list = append(list, tmp...)
	return string(list)
}

func main() {
	fmt.Println(reverseLeftWords("abc", 3))
}
