package _55

import (
	"fmt"
	"sort"
	"testing"
)

func TestCanJump(t *testing.T) {
	for i, tc := range []struct {
		Nums []int
		Ret  bool
	}{
		{[]int{1, 2, 3}, true},
		{[]int{1, 0, 0}, false},
	} {
		if tc.Ret != canJump(tc.Nums) {
			t.Errorf("No.%d result error", i)
		}
		sort.Reverse(sort.IntSlice(tc.Nums))
		fmt.Println()
	}

}

func canJump(nums []int) bool {
	l := len(nums)
	zero := -1

	for i := l - 2; i >= 0; i-- {
		// 已经有0值了
		if zero > 0 {
			// 可以跳过0值
			if i+nums[i] > zero {
				// 当前0值可以忽略
				zero = -1
			}

			continue
		}

		if nums[i] == 0 {
			zero = i
			continue
		}

	}

	return zero < 0
}
