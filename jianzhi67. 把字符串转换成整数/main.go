package main

import (
	"fmt"
	"math"
)

func strToInt1(str string) int {
	sign := 0
	value := 0
	start := false
	mul := func(v int, sign int) int {
		if sign == 0 {
			return v
		}
		return sign * v
	}
	for i := 0; i < len(str); i++ {
		b := str[i]
		switch {
		case b == ' ':
			if start {
				return mul(value, sign)
			}
			continue
		case b == '+':
			if sign != 0 {
				return mul(value, sign)
			}
			sign = 1
			start = true
		case b == '-':
			if sign != 0 {
				return mul(value, sign)
			}
			sign = -1
			start = true
		case '0' <= b && b <= '9':
			a := int(b - '0')
			value = value*10 + a
			if mul(value, sign) >= math.MaxInt32 {
				return math.MaxInt32
			}
			if mul(value, sign) <= math.MinInt32 {
				return math.MinInt32
			}
			start = true
			if sign == 0 {
				sign = 1
			}
		default:
			return mul(value, sign)
		}
	}
	return mul(value, sign)
}

func strToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	// 删空格
	i := 0
	for ; i < len(str); i++ {
		if str[i] != ' ' {
			break
		}
	}
	str = str[i:]
	// 确认正负号
	sign := 0

	if len(str) > 0 {
		if str[0] == '+' {
			sign = 1
		} else if str[0] == '-' {
			sign = -1
		}
	}

	if sign != 0 {
		str = str[1:]
	} else {
		sign = 1
	}

	value := 0
	for i := 0; i < len(str); i++ {
		a := int(str[i] - '0')
		if a < 0 || a > 9 {
			break
		}
		value = value*10 + a
		if value*sign > math.MaxInt32 {
			return math.MaxInt32
		} else if value*sign < math.MinInt32 {
			return math.MinInt32
		}
	}
	value *= sign
	return value
}

func main() {
	// fmt.Println(strToInt("-123"))
	// fmt.Println(strToInt("--123"))
	// fmt.Println(strToInt("       123"))
	// fmt.Println(strToInt("      - 123"))
	// fmt.Println(strToInt("      -123asss"))
	// fmt.Println(strToInt("      -91283472332"))
	// fmt.Println(strToInt("      -1 1"))
	// fmt.Println(strToInt(" 52"))
	// fmt.Println(strToInt("2147483648"))
	// fmt.Println(strToInt("2147483646"))
	// fmt.Println(strToInt("-2147483646"))
	// fmt.Println(strToInt("-2147483649"))
	fmt.Println(strToInt("9223372036854775808"))
}
