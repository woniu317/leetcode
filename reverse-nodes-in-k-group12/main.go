package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if k == 0 || k == 1 || head == nil {
		return head
	}
	var result, t *ListNode
	for {
		if tail, ok := getKSubList(head, k-1); ok {
			temp := tail
			tail = tail.Next
			temp.Next = nil

			reverseListHead := reverseList(head)
			t = appendList(reverseListHead, t)
			if result == nil {
				result = t
			}

			head = tail

		} else {
			t = appendList(tail, t)
			if result == nil {
				result = t
			}
			break
		}

	}
	return result
}

//list1附加list2之后，返回list2的最后一个元素的指针
func appendList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list2 == nil {
		list2 = list1
	} else {
		for list2.Next != nil {
			list2 = list2.Next
		}
		list2.Next = list1
	}
	return list2
}

//unsafe：不做卫术语句
func reverseList(head *ListNode) *ListNode {
	front := head.Next
	head.Next = nil
	for {
		if front == nil {
			return head
		}
		if front.Next == nil {
			front.Next = head
			return front
		}
		temp := front.Next
		front.Next = head
		head = front
		front = temp
	}
}

//返回链表head的第k个元素的指针，若元素个数小于k，则ok=false
func getKSubList(head *ListNode, k int) (tail *ListNode, ok bool) {
	if head == nil {
		return
	}
	h := head
	for i := 0; i < k; i++ {
		head = head.Next
		if head == nil {
			return h, false
		}
	}
	return head, true
}

func main() {
	a1 := ListNode{Val: 1}
	a2 := ListNode{Val: 2}
	a3 := ListNode{Val: 3}
	a4 := ListNode{Val: 4}
	a5 := ListNode{Val: 5}

	a1.Next = &a2
	a2.Next = &a3
	a3.Next = &a4
	a4.Next = &a5

	h := reverseKGroup(&a1, 2)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}
