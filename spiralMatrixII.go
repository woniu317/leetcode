package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	var result [][]int
	for i := 0; i < n; i++ {
		result = append(result, make([]int, n))
	}

	value := 1
	i, j := 0, 0
	xMin, yMin := 0, 0
	xMax, yMax := n-1, n-1
	result[0][0] = value
	for {
		for j < xMax {
			j++
			value++
			result[i][j] = value
		}
		yMin++
		if yMin > yMax {
			break
		}

		for i < yMax {
			i++
			value++
			result[i][j] = value
		}
		xMax--
		if xMin > xMax {
			break
		}

		for j > xMin {
			j--
			value++
			result[i][j] = value
		}
		yMax--
		if yMin > yMax {
			break
		}

		for i > yMin {
			i--
			value++
			result[i][j] = value
		}
		xMin++
		if xMin > xMax {
			break
		}

	}
	return result
}
