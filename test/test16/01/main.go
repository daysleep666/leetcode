package main

func countAndSay(n int) string {
	arr := make([][]byte, n+1)
	arr[1] = []byte{'1'}
	for i := 2; i <= n; i++ {
		arr[i] = cal(arr[i-1])
	}
	return string(arr[n])
}

func cal(list []byte) []byte {
	result := make([]byte, 0, len(list)*2)
	c := 1
	i := 1
	for ; i < len(list); i++ {
		if list[i] == list[i-1] {
			c++
			continue
		}
		result = append(result, byte('0'+c), list[i-1])
		c = 1
	}
	result = append(result, byte('0'+c), list[i-1])
	return result
}
