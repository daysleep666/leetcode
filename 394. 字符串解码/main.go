package main

import (
	"fmt"
	"strconv"
)

// "3[a2[c]]""
func decodeString(s string) string {
	isDigist := func(r rune) bool {
		return '0' <= r && r <= '9'
	}
	stack := make([]rune, 0, len(s))
	for _, v := range s {
		if v != ']' {
			stack = append(stack, v)
			continue
		}
		for i := len(stack) - 1; i >= 0; i-- {
			if stack[i] != '[' {
				continue
			}
			tmp := make([]rune, len(stack)-(i+1))
			copy(tmp, stack[i+1:len(stack)])
			var numStr []rune
			k := i - 1
			for ; k >= 0; k-- {
				if !isDigist(stack[k]) {
					break
				}
				numStr = stack[k:i]
			}
			num, err := strconv.ParseInt(string(numStr), 10, 64)
			if err != nil {
				panic(err)
			}
			stack = stack[:k+1]
			for j := int64(0); j < num; j++ {
				stack = append(stack, tmp...)
			}
			break
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("100[leetcode]"))
}
