package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//先处理头
	h := head
	tail := head.Next
	head.Next = tail.Next
	tail.Next = head
	h = tail

	front, tail := h.Next.Next, h.Next
	for front != nil && front.Next != nil {
		//交换
		tail.Next = front.Next
		front.Next = front.Next.Next
		tail.Next.Next = front

		tail = tail.Next.Next
		front = tail.Next
	}
	return h
}

func main() {

}
