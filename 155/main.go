package main

type MinStack struct {
	data []node
}

type node struct {
	val        int
	currentMin int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{data: []node{}}
}

func (this *MinStack) Push(x int) {
	m := x
	if length := len(this.data); length > 0 {
		m = min(this.data[length-1].currentMin, x)
	}
	this.data = append(this.data, node{
		val:        x,
		currentMin: m,
	})
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (this *MinStack) Pop() {
	if length := len(this.data); length > 0 {
		this.data = this.data[:length-1]
	}
}

func (this *MinStack) Top() int {
	if length := len(this.data); length > 0 {
		return this.data[length-1].val
	}
	panic("")
}

func (this *MinStack) GetMin() int {
	if length := len(this.data); length > 0 {
		return this.data[length-1].currentMin
	}
	panic("")
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
