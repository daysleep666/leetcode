package main

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}
	var i, j int
	push := func() {
		i++
	}
	pop := func() {
		pushed = append(pushed[:i], pushed[i+1:]...)
		if i > 0 {
			i--
		}
		j++
	}
	for i < len(pushed) && j < len(popped) {
		if pushed[i] != popped[j] {
			push()
			continue
		}
		pop()
	}
	return j == len(popped)
}

func main() {

}
