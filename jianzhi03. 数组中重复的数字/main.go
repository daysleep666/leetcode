package main

func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return v
		}
		m[v] = struct{}{}
	}
	return 0
}

func main() {

}
