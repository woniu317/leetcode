package main

import (
	"container/heap"
	"fmt"
)

func main() {
	da := findKthLargest([]int{8, 8, 1, 2, 3, 4, 5, 6, 7, 7}, 3)
	fmt.Println(da)
}

func findKthLargest(data []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i, v := range data {
		if i < k {
			heap.Push(h, v)
		}
	}
	for i := k; i < len(data); i++ {
		//堆顶元素
		top := heap.Pop(h).(int)
		if data[i] > top {
			heap.Push(h, data[i])
		} else {
			heap.Push(h, top)
		}
	}
	return heap.Pop(h).(int)
}

type IntHeap []int

func (intHeap IntHeap) Less(i, j int) bool {
	return intHeap[i] < intHeap[j]
}

func (intHeap IntHeap) Swap(i, j int) {
	intHeap[i], intHeap[j] = intHeap[j], intHeap[i]
}

func (intHeap IntHeap) Len() int {
	return len(intHeap)
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
