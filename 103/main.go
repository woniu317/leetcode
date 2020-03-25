package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	queue := []int{}
	fmt.Println(len(queue), cap(queue))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	queue := []*TreeNode{}
	order := true
	queue = append(queue, root)
	for len(queue) > 0 {
		newQ := []*TreeNode{}
		currentLevelNodes := []int{}
		if order {
			for i := len(queue) - 1; i >= 0; i-- {
				if queue[i] == nil {
					continue
				}
				currentLevelNodes = append(currentLevelNodes, queue[i].Val)
				newQ = append(newQ, queue[i].Left)
				newQ = append(newQ, queue[i].Right)
			}
		} else {
			for i := len(queue) - 1; i >= 0; i-- {
				if queue[i] == nil {
					continue
				}
				currentLevelNodes = append(currentLevelNodes, queue[i].Val)
				newQ = append(newQ, queue[i].Right)
				newQ = append(newQ, queue[i].Left)
			}
		}
		order = !order
		queue = newQ
		if len(currentLevelNodes) != 0 {
			ret = append(ret, currentLevelNodes)
		}
	}
	return ret
}
