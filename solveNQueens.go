package main

import (
	"fmt"
	"reflect"
	//"strings"
)

func main() {
	//fmt.Println(solveNQueens(3))
	va := make([][]int, 2)
	for i, _ := range va {
		va[i] = make([]int, 5)
	}
	for i, v := range va {
		fmt.Println(i, v)
	}
}

var result [][]string
var col []int
