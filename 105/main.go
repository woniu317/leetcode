package main

/**
 * 根据先序遍历和中序遍历建树
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	index := 0
	rootValue := preorder[0]
	for i := range inorder {
		if rootValue == inorder[i] {
			index = i
			break
		}
	}
	root := &TreeNode{Val: rootValue}
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}
