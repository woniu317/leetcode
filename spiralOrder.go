package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	fmt.Println(spiralOrder([][]int{{2, 5}, {8, 4}, {0, -1}}))
	json.Unmarshal()
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	i, j := 0, 0
	xMin, yMin := 0, 0
	xMax, yMax := len(matrix[0])-1, len(matrix)-1
	var result []int
	result = append(result, matrix[0][0])
	for {
		for i < xMax {
			i++
			result = append(result, matrix[j][i])
		}
		yMin++
		if yMin > yMax {
			break
		}

		for j < yMax {
			j++
			result = append(result, matrix[j][i])
		}
		xMax--
		if xMin > xMax {
			break
		}

		for i > xMin {
			i--
			result = append(result, matrix[j][i])
		}
		yMax--
		if yMin > yMax {
			break
		}

		for j > yMin {
			j--
			result = append(result, matrix[j][i])
		}
		xMin++
		if xMin > xMax {
			break
		}
	}
	return result
}
