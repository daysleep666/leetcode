package main

import (
	"fmt"
)

func calculate(s string) int {
	// 1. 去除空格
	// 2. 栈
	// s = strings.TrimSpace(s)
	stack := make([]int, 1)
	isDigist := func(b byte) bool {
		return '0' <= b && b <= '9'
	}
	toNum := func(b byte) int {
		return int(b - '0')
	}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		v := s[i]
		if isDigist(v) {
			stack[len(stack)-1] = stack[len(stack)-1]*10 + toNum(v)
		}
		// fmt.Println("-->", string(v))
		switch v {
		case '+':
			stack = append(stack, -1, 0)
		case '-':
			stack = append(stack, -2, 0)
		case '*':
			i++
			tmp := 0
			for ; i < len(s) && (isDigist(s[i]) || s[i] == ' '); i++ {
				if s[i] == ' ' {
					continue
				}
				tmp = tmp*10 + toNum(s[i])
			}
			stack[len(stack)-1] = stack[len(stack)-1] * tmp
			// fmt.Println("-->", stack)
			i--

		case '/':
			i++
			tmp := 0
			// fmt.Println("-->", stack, string(s[i]))
			for ; i < len(s) && (isDigist(s[i]) || s[i] == ' '); i++ {
				if s[i] == ' ' {
					continue
				}
				tmp = tmp*10 + toNum(s[i])
			}
			stack[len(stack)-1] = stack[len(stack)-1] / tmp
			i--
		default:
			// num = num*10 + toNum(v)
		}
	}
	// fmt.Println(stack)
	for i := 0; i < len(stack)-2; i += 2 {
		a := stack[i]
		s := stack[i+1]
		if s == -1 {
			stack[i+2] = a + stack[i+2]
		} else {
			stack[i+2] = a - stack[i+2]
		}
	}
	return stack[len(stack)-1]
}

func main() {
	// a := byte('9') - '0'
	// fmt.Println(int(a))
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate("3/2"))
	fmt.Println(calculate("3+5/2"))
	fmt.Println(calculate(" 3+5 / 2 "))
}
