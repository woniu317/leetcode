package main

import (
	"container/heap"
)

type ListNode1 struct {
	Val  int
	Next *ListNode1
}
type ListHeap []*ListNode1

func (heap ListHeap) Less(i, j int) bool {
	return heap[i].Val < heap[j].Val
}

func (heap ListHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap ListHeap) Len() int {
	return len(heap)
}

func (heap *ListHeap) Push(x interface{}) {
	*heap = append(*heap, x.(*ListNode1))
}

func (heap *ListHeap) Pop() interface{} {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode1) *ListNode1 {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	h := &ListHeap{}
	heap.Init(h)
	for _, p := range lists {
		if p != nil {
			heap.Push(h, p)
			p = p.Next
		}
	}
	var head *ListNode1
	var point *ListNode1
	for len(*h) > 1 {
		firstElem := heap.Pop(h).(*ListNode1)
		if head == nil {
			head = firstElem
			point = head
		} else {
			point.Next = firstElem
			point = point.Next
		}
		firstElem = firstElem.Next
		if firstElem != nil {
			heap.Push(h, firstElem)
		}
	}
	if len(*h) == 1 {
		point.Next = heap.Pop(h).(*ListNode1)
	}
	return head
}

func main() {
	node1 := ListNode1{Val: 5}
	node2 := ListNode1{Val: 12}
	node1.Next = &node2

	node21 := ListNode1{Val: 2}
	node22 := ListNode1{Val: 22}
	node21.Next = &node22
	lists := make([]*ListNode1, 2)
	lists[0] = &node1
	lists[1] = &node21
	mergeKLists(lists)
}
