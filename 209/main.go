package main

func main() {
	s := 15
	nums := []int{1, 2, 3, 4, 5}
	println(minSubArrayLen(s, nums))
}

func minSubArrayLen(s int, nums []int) int {
	front := 0
	if len(nums) == 0 {
		return 0
	}
	lenth := len(nums)
	curMinLen := lenth + 1
	curSum := 0
	for i := 0; i < lenth; i++ {
		curSum += nums[i]
		if curSum >= s {
			ms := minSub(front, i, s, &nums)
			if ms < curMinLen {
				curMinLen = ms
			}
		}
	}
	if curMinLen == lenth+1 {
		return 0
	}
	return curMinLen
}

func minSub(front, cur, s int, nums *[]int) int {
	curSum := 0
	curLen := 0
	for ; front <= cur; cur-- {
		curLen++
		curSum += (*nums)[cur]
		if curSum >= s {
			break
		}
	}
	return curLen
}
