package main

import "fmt"

func generateParenthesis1(n int) (result []string) {
	if n == 0 {
		return nil
	}
	return dfs("(", 1, 0, n)
}

func dfs(cur string, l, r, max int) (result []string) {
	if r >= max {
		return []string{cur}
	}
	if l < max {
		r1 := dfs("(", l+1, r, max)
		for _, v := range r1 {
			result = append(result, cur+v)
		}
	}
	if l > r {
		r2 := dfs(")", l, r+1, max)
		for _, v := range r2 {
			result = append(result, cur+v)
		}
	}
	return result
}

func generateParenthesis(n int) (result []string) {
	if n == 0 {
		return
	}
	helper(&result, "(", 1, 0, n)
	return
}

func helper(result *[]string, cur string, l, r, max int) {
	if r == max {
		*result = append(*result, cur)
		return
	}
	if l < max {
		helper(result, cur+"(", l+1, r, max)
	}
	if l > r {
		helper(result, cur+")", l, r+1, max)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}
