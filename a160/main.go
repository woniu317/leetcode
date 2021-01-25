package main

import (
	"fmt"
	"sync"
)

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := getListLen(headA)
	lenB := getListLen(headB)
	for lenA > lenB {
		headA = headA.Next
		lenA--
	}
	for lenB > lenA {
		headB = headB.Next
		lenB--
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

func getListLen(head *ListNode) int {
	ret := 0
	for head != nil {
		ret++
		head = head.Next
	}
	return ret
}

func main() {
	headA := &ListNode{Val: 0}
	n1 := &ListNode{Val: 9}
	n2 := &ListNode{Val: 1}
	n3 := &ListNode{Val: 2}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	headA.Next = n1

	headB := &ListNode{Val: 3}
	headB.Next = n3
	printList(headA)
	printList(headB)

	fmt.Println(getIntersectionNode(headA, headB).Val)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go funcName(wg)
}

func funcName(wg sync.WaitGroup) {
	func() {
		defer wg.Done()
	}()
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
