package main

import (
	"fmt"
	"sort"
)

func threeSum(data []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(data)
	if len(data) == 0 {
		return result
	}
	for k := 0; k < len(data); k++ {
		if k > 0 && data[k] == data[k-1] {
			continue
		}
		tmpTarget := target - data[k]
		twoResult := twoNum(data, k+1, tmpTarget)
		for _, da := range twoResult {
			res := []int{data[k], da[0], da[1]}
			result = append(result, res)
		}
	}

	return result
}

func twoNum(data []int, startPosition int, target int) [][]int {
	result := make([][]int, 0)
	i, j := startPosition, len(data)-1
	for i < j {
		if data[i]+data[j] == target {
			res := []int{data[i], data[j]}
			result = append(result, res)
			for i < j && data[i] == data[i+1] {
				i++
			}
			for i < j && data[j] == data[j-1] {
				j--
			}
			i++
			j--
		} else if data[i]+data[j] < target {
			i++
		} else {
			j--
		}
	}
	return result
}

func main() {
	tn := threeSum([]int{-1, -1, -2, -2, -3, 0, 1, 3, 3, 3, 5, 8}, 0)
	for _, da := range tn {
		fmt.Println(da)
	}
}
