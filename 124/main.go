package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return math.MinInt32
	}
	currentMax := math.MinInt32
	maxPathSumSub(root, &currentMax)
	return currentMax
}
func maxPathSumSub(root *TreeNode, currentMax *int) int {
	if root == nil {
		return math.MinInt32
	}
	leftMax := maxPathSumSub(root.Left, currentMax)
	rightMax := maxPathSumSub(root.Right, currentMax)
	m := max(leftMax, rightMax)
	m = max(root.Val, m+root.Val)
	t := max(m, root.Val+leftMax+rightMax)
	*currentMax = max(*currentMax, t)
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max(9, 9+math.MinInt32+math.MinInt32))
}
