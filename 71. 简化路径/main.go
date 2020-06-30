package main

import "fmt"

// 栈
func simplifyPath(path string) string {
	stack := make([]string, 0, 10)
	pop := func() {
		if len(stack) == 0 {
			return
		}
		stack = stack[:len(stack)-1]
	}
	push := func(v string) {
		stack = append(stack, v)
	}
	for i := 0; i < len(path); {
		if path[i] == '/' {
			i++
			continue
		}

		tmp := []byte{}
		for ; i < len(path); i++ {
			if path[i] == '/' {
				break
			}
			tmp = append(tmp, path[i])
		}
		t := string(tmp)
		if t == "." {
			continue
		} else if t == ".." {
			pop()
			continue
		}
		push(t)
	}

	if len(stack) == 0 {
		return "/"
	}
	result := ""
	for _, v := range stack {
		result = result + "/" + v
	}
	return result
}

func main() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/a/.."))
}
