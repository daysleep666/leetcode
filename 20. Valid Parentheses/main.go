package main

import "fmt"

var m = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	list := make([]rune, 0, len(s))
	for _, key := range s {
		_, ok := m[key]
		if ok { // 左侧
			list = append(list, key)
		} else { // 右侧
			if len(list) == 0 {
				return false
			}
			x := list[len(list)-1]
			if m[x] != key {
				return false
			}
			list = list[:len(list)-1]
		}
	}
	return len(list) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("()}"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid(""))
}

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.

// Example 1:

// Input: "()"
// Output: true
// Example 2:

// Input: "()[]{}"
// Output: true
// Example 3:

// Input: "(]"
// Output: false
// Example 4:

// Input: "([)]"
// Output: false
// Example 5:

// Input: "{[]}"
// Output: true
