package main

import "fmt"

func main() {
	// fmt.Println(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))
	// fmt.Println(kthSmallest([][]int{{1, 2}, {1, 3}}, 4))
	fmt.Println(kthSmallest([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
}

func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	first := make([]int, len(matrix))
	index := 0
	for k > 0 {
		index = findMin(matrix, first)
		first[index]++
		fmt.Println(index)
		k--
	}
	return matrix[index][first[index]-1]
}

func findMin(matrix [][]int, index []int) int {
	rowLen := len(matrix)
	colLen := len(matrix[0])
	tmp := matrix[rowLen-1][colLen-1] + 1
	r := 0
	for i := 0; i < len(matrix); i++ {
		if index[i] == colLen {
			continue
		}
		if matrix[i][index[i]] < tmp {
			tmp = matrix[i][index[i]]
			r = i
		}
	}
	return r
}

// [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]
// 20
// 21
