package main

import "fmt"

func main() {
	g := grayCode(2)
	fmt.Print(g)
}

func grayCode(n int) []int {
	ret := make([]int, 1, 1<<n)
	for i := 0; i < n; i++ {
		t := 1 << i
		for j := len(ret) - 1; j >= 0; j-- {
			ret = append(ret, ret[j]+t)
		}
	}
	return ret
}
