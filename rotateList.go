package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	temHead := head
	length := 0
	for temHead != nil {
		temHead = temHead.Next
		length++
	}
	k = k % length
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	slow, fast := head, head
	for i := 0; i <= k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fast = slow.Next
	slow.Next = nil
	slow = fast
	for fast.Next != nil {
		fast = fast.Next
	}
	fast.Next = head
	return slow
}

func main() {
	rotateRight(nil, 1)
}
