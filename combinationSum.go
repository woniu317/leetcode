package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := new([][]int)
	curRes := new([]int)
	firstIndex := 0
	dfs(res, curRes, firstIndex, candidates, target)
	return *res
}

func dfs(res *[][]int, curRes *[]int, firstIndex int, candidates []int, target int) {
	if target < 0 {
		return
	} else if target == 0 {
		oneResult := make([]int, len(*curRes))
		copy(oneResult, *curRes)
		*res = append(*res, oneResult)
		return
	}
	for i := firstIndex; i < len(candidates) && candidates[i] <= target; i++ {
		*curRes = append(*curRes, candidates[i])
		dfs(res, curRes, i, candidates, target-candidates[i])
		*curRes = (*curRes)[0 : len(*curRes)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := new([][]int)
	curRes := new([]int)
	firstIndex := 0
	dfs2(res, curRes, firstIndex, candidates, target)
	return *res
}

func dfs2(res *[][]int, curRes *[]int, firstIndex int, candidates []int, target int) {
	if target < 0 {
		return
	} else if target == 0 {
		oneResult := make([]int, len(*curRes))
		copy(oneResult, *curRes)
		*res = append(*res, oneResult)
		return
	}
	lastCandidates := 0
	for i := firstIndex; i < len(candidates) && candidates[i] <= target; i++ {
		if candidates[i] == lastCandidates {
			continue
		}
		lastCandidates = candidates[i]
		*curRes = append(*curRes, candidates[i])
		dfs2(res, curRes, i+1, candidates, target-candidates[i])
		*curRes = (*curRes)[0 : len(*curRes)-1]
	}
}

func main() {
	//fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	//fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
