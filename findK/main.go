package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func main() {
	da := findK([]int{8, 8, 8, 1, 2, 3, 4, 5, 6, 7, 7}, 3)
	for _, v := range da {
		println(v)
	}
}

func findK(data []int, k int) []int {
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
		if data[i] < top {
			heap.Push(h, data[i])
		} else {
			heap.Push(h, top)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).(int)
	}
	return res
}
