package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
func jump(nums []int) int {
	jumps, curEnd, curFast := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		curFast = max(curFast, nums[i]+i)
		if i == curEnd {
			jumps++
			curEnd = curFast
		}
	}
	return jumps
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//public int jump(int[] A) {
//int jumps = 0, curEnd = 0, curFarthest = 0;
//for (int i = 0; i < A.length - 1; i++) {
//curFarthest = Math.max(curFarthest, i + A[i]);
//if (i == curEnd) {
//jumps++;
//curEnd = curFarthest;
//}
//}
//return jumps;
//}
