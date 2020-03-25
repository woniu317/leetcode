package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	dataStack := []int{}
	for _, token := range tokens {
		l := len(dataStack)
		switch token {
		case "+":
			dataStack = append(dataStack[:l-2], dataStack[l-2]+dataStack[l-1])
		case "-":
			dataStack = append(dataStack[:l-2], dataStack[l-2]-dataStack[l-1])
		case "*":
			dataStack = append(dataStack[:l-2], dataStack[l-2]*dataStack[l-1])
		case "/":
			dataStack = append(dataStack[:l-2], dataStack[l-2]/dataStack[l-1])
		default:
			tmp, _ := strconv.Atoi(token)
			dataStack = append(dataStack, tmp)
		}
	}
	return dataStack[0]
}

func main() {
	rpn := convertToRPN([]string{"1", "+", "2", "+", "3", "*", "4", "/", "2"})
	fmt.Println(rpn)
	res := evalRPN(rpn)
	fmt.Println(res)
}

func convertToRPN(tokens []string) []string {
	ret := []string{}

	stack1 := []string{}
	stack1 = append(stack1, "#")

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			topOperator := stack1[len(stack1)-1]
			for !greater(token, topOperator) {
				ret = append(ret, topOperator)
				stack1 = stack1[:len(stack1)-1]
				topOperator = stack1[len(stack1)-1]
			}
			stack1 = append(stack1, token)
		default:
			ret = append(ret, token)
		}
	}
	for i := len(stack1) - 1; i > 0; i-- {
		ret = append(ret, stack1[i])
	}

	return ret
}

func greater(operator1 string, operator2 string) bool {
	return pririoty(operator1) > pririoty(operator2)
}

func pririoty(operator string) int {
	switch operator {
	case "#":
		return 0
	case "+", "-":
		return 1
	default:
		return 2
	}
}
