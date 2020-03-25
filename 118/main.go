package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	ret := [][]int{}
	if numRows == 0 {
		return ret
	}
	ret = append(ret, []int{1})
	if numRows == 1 {
		return ret
	}
	ret = append(ret, []int{1, 1})
	if numRows == 2 {
		return ret
	}
	for i := 3; i <= numRows; i++ {
		tmp := []int{1}
		for j := 1; j < i-1; j++ {
			v := ret[i-2][j-1] + ret[i-2][j]
			tmp = append(tmp, v)
		}
		tmp = append(tmp, 1)
		ret = append(ret, tmp)
	}
	return ret
}
