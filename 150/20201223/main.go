package main

import (
	"fmt"
	"strconv"
)

func main() {
	token := tokenizer("-1+(2-3)*4*(-5)+10/(2+3)")
	fmt.Println(token)
	rpn := convertToRPN(token)
	fmt.Println(rpn)
	v := valueRPN(rpn)
	fmt.Println(v)
}

func valueRPN(token []string) int {
	var value []int
	for _, val := range token {
		l := len(value)
		switch val {
		case "+":
			value = append(value[:l-2], value[l-2]+value[l-1])
		case "-":
			value = append(value[:l-2], value[l-2]-value[l-1])
		case "*":
			value = append(value[:l-2], value[l-2]*value[l-1])
		case "/":
			value = append(value[:l-2], value[l-2]/value[l-1])
		default:
			v, _ := strconv.Atoi(val)
			value = append(value, v)
		}
	}
	return value[0]
}

func tokenizer(expression string) []string {
	var ret []string
	start := -1

	for i, b := range expression {
		switch b {
		case '-':
			if start != -1 {
				ret = append(ret, expression[start:i])
				ret = append(ret, string(b))
				start = -1
			} else {
				start = i
			}
		case '+', '(', '*', '/', ')':
			if start != -1 {
				ret = append(ret, expression[start:i])
			}
			ret = append(ret, string(b))
			start = -1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if start == -1 {
				start = i
			}
		default:
			panic("unknown notation")
		}
	}

	if start != -1 {
		ret = append(ret, expression[start:])
	}

	return ret
}

func convertToRPN(token []string) []string {
	ret := make([]string, 0, len(token))
	notation := []string{"#"}

	for _, t := range token {
		switch t {
		case "+", "-", "*", "/":
			top := notation[len(notation)-1]
			for greaterOrEqual(top, t) {
				ret = append(ret, top)
				notation = notation[:len(notation)-1]
				top = notation[len(notation)-1]
			}
			notation = append(notation, t)
		case "(":
			notation = append(notation, t)
		case ")":
			top := notation[len(notation)-1]
			for top != "(" {
				ret = append(ret, top)
				notation = notation[:len(notation)-1]
				top = notation[len(notation)-1]
			}
			notation = notation[:len(notation)-1]
		default:
			ret = append(ret, t)
		}
	}

	for i := len(notation) - 1; i > 0; i-- {
		ret = append(ret, notation[i])
	}

	return ret
}

func greaterOrEqual(operator1, operator2 string) bool {
	return priority(operator1) >= priority(operator2)
}

func priority(operator string) int {
	switch operator {
	case "#":
		return 0
	case "(":
		return 10
	case ")":
		return 31
	case "+", "-":
		return 20
	case "*", "/":
		return 30
	default:
		panic("unknown operator")
	}
}
