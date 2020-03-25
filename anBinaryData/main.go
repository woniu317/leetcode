package main

import "fmt"

func parse(n int) []int {
	res := []int{}
	i := 0
	for n > 0 {
		if n != n/2*2 {
			res = append(res, i)
		}
		n = n / 2
		i++
	}
	return res
}

func main() {
	fmt.Println(parse(33))
}
