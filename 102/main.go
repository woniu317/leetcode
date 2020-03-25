package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder1(root *TreeNode) [][]int {
	qr := make([]*TreeNode, 0)
	qf := make([]int, 0)
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	r := root
	qr = append(qr, r)
	qf = append(qf, 1)
	lastRes := make([]int, 0)
	lastRes = append(lastRes, r.Val)
	for len(qr) != 0 {
		v := qr[0]
		f := qf[0]
		if len(qf) == 1 || f != qf[1] {
			if len(lastRes) != 0 {
				tmp := make([]int, len(lastRes))
				copy(tmp, lastRes)
				res = append(res, tmp)
				lastRes = lastRes[:0]
			}
		}
		qr = qr[1:]
		qf = qf[1:]
		if v.Left != nil {
			qr = append(qr, v.Left)
			lastRes = append(lastRes, v.Left.Val)
			qf = append(lastRes, f+1)
		}
		if v.Right != nil {
			qr = append(qr, v.Right)
			lastRes = append(lastRes, v.Right.Val)
			qf = append(lastRes, f+1)
		}
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	q := []*TreeNode{}
	q = append(q, root)
	newQ := []*TreeNode{}
	for len(q) > 0 {
		currentLevelNodes := []int{}
		for _, r := range q {
			if r == nil {
				continue
			}
			currentLevelNodes = append(currentLevelNodes, r.Val)
			newQ = append(newQ, r.Left)
			newQ = append(newQ, r.Right)
		}
		if len(currentLevelNodes) != 0 {
			ret = append(ret, currentLevelNodes)
		}
		q, newQ = newQ, q
		newQ = newQ[:0]
	}
	return ret
}
