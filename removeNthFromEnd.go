package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//能通过leetcode验证，但是实际工程里面肯定要判断n的大小，因为若不判断，则当n>len(head)或head=nil时，就会有异常
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fasterPoint, slowPoint := head, head
	for i := 0; i < n; i++ {
		fasterPoint = fasterPoint.Next
	}
	if fasterPoint == nil {
		return head.Next
	}
	for fasterPoint.Next != nil {
		fasterPoint = fasterPoint.Next
		slowPoint = slowPoint.Next
	}
	slowPoint.Next = slowPoint.Next.Next
	return head
}

func main() {
	removeNthFromEnd(nil, 1)
}
