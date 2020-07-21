package main

// 双指针
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	// 是否是字母
	isAlpha := func(a byte) bool {
		return ('a' <= a && a <= 'z') || ('A' <= a && a <= 'Z')
	}

	// 是否是数字
	isDigist := func(a byte) bool {
		return '0' <= a && a <= '9'
	}

	toLower := func(a byte) byte {
		if 'A' <= a && a <= 'Z' {
			return a - 'A' + 'a'
		}
		return a
	}

	// start,end
	for start, end := 0, len(s)-1; start < end; {
		a, b := s[start], s[end]
		if !isAlpha(a) && !isDigist(a) {
			start++
			continue
		}
		if !isAlpha(b) && !isDigist(b) {
			end--
			continue
		}
		a = toLower(a)
		b = toLower(b)
		if a != b {
			return false
		}
		start++
		end--
	}
	return true
}
