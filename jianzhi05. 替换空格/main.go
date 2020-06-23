package main

func replaceSpace(s string) string {
	r := make([]rune, 0, len(s)*3)
	for _, v := range s {
		if v == ' ' {
			r = append(r, '%')
			r = append(r, '2')
			r = append(r, '0')
			continue
		}
		r = append(r, v)
	}
	return string(r)
}

func main() {

}
