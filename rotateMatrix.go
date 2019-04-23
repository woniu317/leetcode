package main

import "fmt"

func main() {
	//matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	matrix := [][]int{{2, 29, 20, 26, 16, 28}, {12, 27, 9, 25, 13, 21}, {32, 33, 32, 2, 28, 14}, {13, 14, 32, 27, 22, 26}, {33, 1, 20, 7, 21, 7}, {4, 24, 1, 6, 32, 34}}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	length := len(matrix)
	for flow := 0; flow < length/2; flow += 1 {
		curLen := length - 2*flow
		corner := curLen - 1 + flow
		for i := flow; i < curLen-1+flow; i++ {
			t := length - i - 1
			matrix[flow][i], matrix[i][corner], matrix[corner][t], matrix[t][flow] =
				matrix[t][flow], matrix[flow][i], matrix[i][corner], matrix[corner][t]
		}
	}
}
