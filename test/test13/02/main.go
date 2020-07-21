package main

// 栈
// 双指针
func removeOuterParentheses(S string) string {
	result := make([]byte, 0, len(S))
	start, end := 0, 0
	left, right := 0, 0
	arr := []byte(S)
	for ; end < len(arr); end++ {
		if arr[end] == '(' {
			left++
		} else {
			right++
		}
		if left != right {
			continue
		}
		result = append(result, arr[start+1:end]...)
		start = end + 1
	}
	return string(result)
}
