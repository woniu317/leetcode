package main

type ListNode struct {
	Next *ListNode
}

func main() {
	head := &ListNode{}
	head.Next = &ListNode{}
	println(listLenth(head))
}

func listLenth(head *ListNode) int {
	lenth := 0
	for head != nil {
		lenth++
		head = head.Next
	}
	return lenth
}
