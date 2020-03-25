package main

/**
 * 有序链表构建二叉平衡树
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var left *ListNode = nil
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		left = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if left == nil {
		return &TreeNode{Val: head.Val}
	}
	left.Next = nil
	root := &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(slow.Next),
	}
	return root
}
