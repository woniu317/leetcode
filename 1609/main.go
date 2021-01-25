package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	q := []*TreeNode{root}
	var leftFirst bool

	for len(q) > 0 {
		var t []*TreeNode
		for i := range q {
			if q[i] == nil {
				continue
			}

			if i > 0 && q[i].Val < q[i-1].Val {
				return false
			}

			if leftFirst {
				if q[i].Val%2 != 0 {
					return false
				}

				t = append(t, q[i].Left, q[i].Right)
				continue
			}

			if q[i].Val%2 == 0 {
				return false
			}
			t = append(t, q[i].Right, q[i].Left)
		}

		leftFirst = !leftFirst
		q = t
	}

	return true
}

func main() {
}

func minSubArrayLen(s int, nums []int) int {
	var start, sum int
	ret := math.MaxInt32

	for i := range nums {
		sum += nums[i]
		for sum > s {
			ret = min(ret, i-start+1)
			start++
		}
	}

	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
