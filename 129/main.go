package main

import "fmt"

func main() {
	fmt.Println(prevPermOpt1([]int{2, 9, 1, 2, 7, 8}))
}
func prevPermOpt1(A []int) []int {
	i := len(A) - 2
	// 找到逆向破坏降序结构的第一个数字
	for ; i >= 0; i-- {
		if A[i] > A[i+1] {
			break
		}
	}
	if i == -1 {
		return A
	}

	j := len(A) - 1
	// 逆向找到i后面比i小的最大值
	for ; j > i; j-- {
		if A[j] < A[i] {
			break
		}
	}
	// 找到位数最高的比i小的最大值
	for ; j-1 > i; j-- {
		if A[j] != A[j-1] {
			break
		}
	}
	A[i], A[j] = A[j], A[i]
	return A
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ant int
	dfs(root, 0, &ant)
	return ant
}

func dfs(root *TreeNode, ret int, ant *int) {
	ret = ret*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*ant += ret
		return
	}

	if root.Left != nil {
		dfs(root.Left, ret, ant)
	}

	if root.Right != nil {
		dfs(root.Right, ret, ant)
	}
}
