package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	path := []int{}
	dfs(root, 1, &path)
	return path
}

func dfs(root *TreeNode, level int, path *[]int) {
	if root == nil {
		return
	}
	if level > len(*path) {
		*path = append(*path, root.Val)
	}
	if root.Right != nil {
		dfs(root.Right, level+1, path)
	}
	if root.Left != nil {
		dfs(root.Left, level+1, path)
	}
}
