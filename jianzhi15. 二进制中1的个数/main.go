package main

func hammingWeight(num uint32) int {
	c := uint32(0)
	for num != 0 {
		c += num & 1
		num = num >> 1
	}
	return int(c)
}

func main() {

}
