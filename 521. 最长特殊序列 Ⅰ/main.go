package main

func findLUSlength(a string, b string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if len(a) != len(b) {
		return max(len(a), len(b))
	}
	if a != b {
		return len(a)
	}
	return -1
}

func main() {

}
