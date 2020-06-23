package main

func printNumbers(n int) []int {
	max:=1
	for i:=0;i<n;i++{
		max*=10
	}
	arr:=make([]int, max-1)
	for i:=0;i<max-1;i++{
		arr[i]=i+1
	}
	return arr
}

func main() {

}


