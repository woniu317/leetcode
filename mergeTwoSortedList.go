package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	p1, p2 := l1, l2
	if l1.Val > l2.Val {
		head = l2
		p2 = p2.Next
	} else {
		head = l1
		p1 = p1.Next
	}
	p := head
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
			p = p.Next
			continue
		}
		p.Next = p1
		p1 = p1.Next
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	} else {
		p.Next = p2
	}
	return head
}
