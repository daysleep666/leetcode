package main

func isStraight(nums []int) bool {
	m := make(map[int]bool, len(nums))
	min := 14
	zero := 0
	for _, v := range nums {
		m[v] = true
		if min > v && v != 0 {
			min = v
		}
		if v == 0 {
			zero++
		}
	}
	if min == 14 {
		return true
	}
	for i := min + 1; i < min+5; i++ {
		if ok := m[i]; !ok {
			zero--
			if zero == -1 {
				return false
			}
		}
	}
	return true
}

func main() {

}
