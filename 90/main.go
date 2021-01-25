package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	sl := make([]int, 0, len(nums))
	dfs(nums, &ret, sl)
	return ret
}

func dfs(nums []int, ret *[][]int, sl []int) {
	dst := make([]int, len(sl))
	copy(dst, sl)
	*ret = append(*ret, dst)

	for i, n := range nums {
		if i > 0 && n == nums[i-1] {
			continue
		}
		sl = append(sl, n)
		dfs(nums[i+1:], ret, sl)
		sl = sl[:len(sl)-1]
	}
}
