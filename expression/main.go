package main

import (
	"fmt"
)

func main() {
	fmt.Println(expValue("1+3*4/2-12"))
}

func explanExpress(str string) ([]int, []string) {
	valueSlice := make([]int, 0)
	sh := make([]string, 0)

	for i := 0; i < len(str); {
		if '0' <= str[i] && '9' >= str[i] {
			tmp := 0
			for ; i < len(str); i++ {
				if '0' <= str[i] && '9' >= str[i] {
					tmp *= 10
					tmp += int(str[i] - '0')
				} else {
					break
				}
			}
			valueSlice = append(valueSlice, tmp)
		} else {
			sh = append(sh, string(str[i]))
			i++
		}
	}
	return valueSlice, sh
}

func expValue(str string) int {
	datas, operators := explanExpress(str)
	dh := &IntStack{}
	oh := &stringStack{}
	dataIndex := 0
	operatorIndex := 0
	//heap.Push(dh, datas[0])
	for {
		//取数字
		if dataIndex >= len(datas) {
			break
		}
		tmpValue := datas[dataIndex]
		dataIndex++
		dh.Push(tmpValue)
		if operatorIndex >= len(operators) {
			break
		}
		tmpOperator := operators[operatorIndex]
		operatorIndex++

		if oh.IsEmpty() || great(tmpOperator, oh.Top()) {
			oh.Push(tmpOperator)
		} else {
			prossess(dh, oh, tmpOperator)
			oh.Push(tmpOperator)
		}

	}
	prossess(dh, oh, "#")
	return dh.Pop()
}

func great(s string, s2 string) bool {
	p1 := getPriority(s)
	p2 := getPriority(s2)
	return p1 > p2
}

func prossess(dh *IntStack, oh *stringStack, newOp string) {
	for !oh.IsEmpty() && !great(newOp, oh.Top()) {
		d1 := dh.Pop()
		d2 := dh.Pop()
		operator := oh.Pop()
		res := cal(d2, d1, operator)
		dh.Push(res)
	}
}

func getPriority(s string) int {
	switch s {
	case "#":
		return 0
	case "+":
		return 1
	case "-":
		return 1
	default:
		return 2
	}
}

func cal(i int, i2 int, s string) int {
	switch s {
	case "+":
		return i + i2
	case "-":
		return i - i2
	case "*":
		return i * i2
	default:
		return i / i2
	}
}

func less(oper1, oper2 string) bool {
	return true
}

type IntStack []int

func (h *IntStack) Push(x int) {
	*h = append(*h, x)
}

func (h *IntStack) Pop() int {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *IntStack) Top() interface{} {
	return (*h)[len(*h)-1]
}

func (h *IntStack) IsEmpty() bool {
	return len(*h) == 0
}

type stringStack []string

func (h *stringStack) Push(x string) {
	*h = append(*h, x)
}

func (h *stringStack) Pop() string {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *stringStack) Top() string {
	return (*h)[len(*h)-1]
}

func (h *stringStack) IsEmpty() bool {
	return len(*h) == 0
}
