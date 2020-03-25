package main

/**
 * 根据后序遍历和中序遍历建树
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	rootV := postorder[len(postorder)-1]
	index := 0
	for i, v := range inorder {
		if rootV == v {
			index = i
			break
		}
	}
	root := &TreeNode{Val: rootV}
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
