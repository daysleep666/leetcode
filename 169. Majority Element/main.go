package main

func majorityElement(nums []int) int {
	var a, b int
	for _, v := range nums {
		if b == 0 {
			a = v
			b++
		} else {
			if v == a {
				b++
			} else {
				b--
			}
		}
	}
	return a
}

func main() {

}
