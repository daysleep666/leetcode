package main

import "fmt"

func reverseWords(s string) string {
	result := make([]byte, 0, len(s))
	end := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			for j := end - 1; j > i; j-- {
				if s[j] == ' ' {
					end--
				}
			}
			if i == end-1 {
				continue
			}
			result = append(result, s[i+1:end]...)
			result = append(result, ' ')
			end = i
			continue
		}
	}
	for j := end - 1; j >= 0; j-- {
		if s[j] == ' ' {
			end--
		}
	}
	result = append(result, s[0:end]...)
	i := len(result) - 1
	for ; i >= 0; i-- {
		if result[i] != ' ' {
			break
		}
	}
	result = result[:i+1]
	return string(result)
}

func main() {
	fmt.Println(fmt.Sprintf("%q\n", reverseWords("the    sky is blue")))
	fmt.Println(fmt.Sprintf("%q\n", reverseWords("  abc  ")))
	fmt.Println(fmt.Sprintf("%q\n", reverseWords("  hello world!  ")))

}
