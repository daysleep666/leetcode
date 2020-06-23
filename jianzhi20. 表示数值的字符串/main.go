package main

import "fmt"

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	// 去空格
	i := 0
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	s = s[i:]
	i = len(s) - 1
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}
	s = s[:i+1]
	if len(s) == 0 {
		return false
	}
	// 去除符号
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	if len(s) == 0 {
		return false
	}

	dotCount, eCount := 0, 0
	if s[0] == '.' {
		s = s[1:]
		dotCount = 1
	}
	if len(s) == 0 {
		return false
	}
	if s[0] == 'e' || s[len(s)-1] == 'e' {
		return false
	}
	hasE := false
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '.':
			if dotCount != 0 || hasE {
				return false
			}
			dotCount++
		case s[i] == 'e':
			if eCount != 0 {
				return false
			}
			eCount++
			hasE = true
		case '0' <= s[i] && s[i] <= '9':

		default:
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isNumber("   "))
	fmt.Println(isNumber("1   "))
	fmt.Println(isNumber("1e1"))
	fmt.Println(isNumber(".1"))
	fmt.Println(isNumber(".."))
	fmt.Println(isNumber("."))
	fmt.Println(isNumber("6e6.5"))
}
