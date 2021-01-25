package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// [1,4,3,2,5,2]
	h := &ListNode{Val: 1}
	l := h
	l.Next = &ListNode{Val: 4}
	l = l.Next
	l.Next = &ListNode{Val: 3}
	l = l.Next
	l.Next = &ListNode{Val: 2}
	l = l.Next
	l.Next = &ListNode{Val: 5}
	l = l.Next
	l.Next = &ListNode{Val: 2}
	a := partition(h, 3)
	fmt.Println(a)
}

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	large := &ListNode{}
	cs, cl := small, large
	for head != nil {
		if head.Val < x {
			cs.Next = head
			cs = head
		} else {
			cl.Next = head
			cl = head
		}
		head = head.Next
	}
	cl.Next = nil
	cs.Next = large.Next
	return small.Next
}
