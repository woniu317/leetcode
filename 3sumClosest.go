package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(data []int, target int) int {
	sort.Ints(data)
	if len(data) == 0 {
		return 0
	}
	closestSum := math.MaxInt64
	closestSumDistance := math.MaxInt64
	for k := 0; k < len(data)-2; k++ {
		if k > 0 && data[k] == data[k-1] {
			continue
		}
		tmpTarget := target - data[k]
		twoResult := twoClosetNum(data, k+1, tmpTarget)
		tempDistance := min3(int(math.Abs(float64(twoResult+data[k]-target))), closestSumDistance)
		if tempDistance != closestSumDistance {
			closestSumDistance = tempDistance
			closestSum = twoResult + data[k]
		}
	}

	return closestSum
}

func min3(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func twoClosetNum(data []int, startPosition int, target int) int {
	closetNumDistance := math.MaxInt64
	closetNum := math.MaxInt64
	i, j := startPosition, len(data)-1

	for i < j {
		if data[i]+data[j] == target {
			return target
		}
		temp := min3(closetNumDistance, int(math.Abs(float64(target-(data[i]+data[j])))))
		if temp != closetNumDistance {
			closetNumDistance = temp
			closetNum = data[i] + data[j]
		}
		if data[i]+data[j] < target {
			i++
		} else {
			j--
		}
	}
	return closetNum
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 1, 2}, 3))
}
