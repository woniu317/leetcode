package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
	fmt.Println(longestConsecutive2(nums))
	fmt.Println(longestConsecutive4(nums))
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	maxLen := 0
	preElem := nums[0]
	tmpLen := 0
	for i := 1; i < len(nums); i++ {

		if nums[i]-preElem == 1 {
			tmpLen++
			preElem = nums[i]
			if tmpLen > maxLen {
				maxLen = tmpLen
			}
		} else if nums[i]-preElem == 0 { //注意有重复的元素
			preElem = nums[i]
		} else {
			preElem = nums[i]
			tmpLen = 0
		}
	}
	return maxLen + 1
}

func longestConsecutive2(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	visited := make(map[int]int)
	for _, v := range nums {
		if visited[v] > 0 {
			continue
		}
		visited[v] = 1
		left, right := visited[v-1], visited[v+1]
		if left > 0 && right > 0 {
			visited[v-left] += right + 1
			visited[v+right] = visited[v-left]
			continue
		}
		if left > 0 && right == 0 {
			visited[v-left] += 1
			visited[v] = left + 1
			continue
		}
		if right > 0 && left == 0 {
			visited[v] = right + 1
			visited[v+right] = right + 1
		}
	}
	maxCount := 0
	for _, v := range visited {
		if v > maxCount {
			maxCount = v
		}
	}
	return maxCount
}

func longestConsecutive4(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	// 一方面，visited用于记录数是否出现过(由于前面set已去重，所以这里并未其效果)
	// 另一方面，visited的值为当前这个数所在的连续序列的新长度
	visited := make(map[int]int)

	// 伪动态规划？
	left, right := 0, 0 // 就是两个临时变量
	for i := 0; i < n; i++ {

		if visited[nums[i]] > 0 {
			continue
		} // 访问过了
		visited[nums[i]] = 1
		// 下面这里两个代码段处理了三种情况：
		// 2 3 4 5(cur) 6 7 8
		// 这里要清楚的是只要是visited[k-1]或者visited[k+1] >0
		// 就意味着他们是连续的，中间插不了数了
		// 我们可以在每一次更新的时候更新序列两端记录的序列长度值（只有序列两端的数，其映射的值才是序列长度值）

		left, right = visited[nums[i]-1], visited[nums[i]+1] // 这是因为k-visited[k-1]有可能=k-1，从而导致其原本值被修改
		if left > 0 && right > 0 {
			visited[nums[i]-left] += right + 1             // 序列起点
			visited[nums[i]+right] = visited[nums[i]-left] // 序列终点处记录的序列长度更新
			continue
		}
		if left > 0 && right == 0 {
			visited[nums[i]] += left // k成为新终点
			visited[nums[i]-left] = visited[nums[i]]
			continue
		}
		if left == 0 && right > 0 {
			visited[nums[i]] += right                 // k成为新起点
			visited[nums[i]+right] = visited[nums[i]] // 序列终点处记录的序列长度更新
			continue
		}
	}

	// 遍历visited，得最大值
	maxcount := 0
	for _, v := range visited {
		if v > maxcount {
			maxcount = v
		}
	}

	return maxcount
}
