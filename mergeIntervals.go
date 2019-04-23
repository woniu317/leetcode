package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 2}, {0, 3}, {4, 5}}))
}

func merge(intervals [][]int) [][]int {

	var result [][]int
	if len(intervals) == 0 {
		return result
	}
	sort.Sort(Intervals(intervals))
	result = append(result, intervals[0])
	for _, v := range intervals {
		top := result[len(result)-1]
		if v[0] <= top[1] {
			top[1] = max(top[1], v[1])
		} else {
			result = append(result, v)
		}
	}
	return result
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	if len(intervals) == 0 {
		result = append(result, newInterval)
		return result
	}
	needCompare := false
	result = append(result, intervals[0])
	for i, v := range intervals {
		if i == 0 {
			continue
		}
		if result[len(result)-1][1] >= newInterval[0] {
			result[len(result)-1][1] = max(result[len(result)-1][1], newInterval[1])
			needCompare = true
		} else if !needCompare {
			result = append(result, v)
		} else {
			if result[len(result)-1][1] >= v[0] {
				result[len(result)-1][1] = max(result[len(result)-1][1], v[1])
			} else {
				result = append(result, v)
				needCompare = false
			}
		}
	}

	return result
}
func insert1(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Intervals [][]int

func (p Intervals) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Intervals) Len() int {
	return len(p)
}

func (p Intervals) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}
