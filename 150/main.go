package main

import (
	"fmt"
	"strconv"
)

/**
 * 逆波兰式求值
 */
func evalRPN(tokens []string) int {
	ds := make([]int, 0)
	for _, val := range tokens {
		l := len(ds)
		switch val {
		case "+":
			ds = append(ds[:l-2], ds[l-2]+ds[l-1])
		case "-":
			ds = append(ds[:l-2], ds[l-2]-ds[l-1])
		case "*":
			ds = append(ds[:l-2], ds[l-2]*ds[l-1])
		case "/":
			ds = append(ds[:l-2], ds[l-2]/ds[l-1])
		default:
			v, _ := strconv.Atoi(val)
			ds = append(ds, v)
		}
	}
	return ds[0]
}

func main() {
	rpn := convertToRPN([]string{"1", "+", "2", "+", "3", "*", "4"})
	fmt.Println(rpn)
	res := evalRPN(rpn)
	fmt.Println(res)
}

func convertToRPN(tokens []string) []string {
	s1 := make([]string, 0)
	var s2 = []string{"#"}
	for _, val := range tokens {
		switch val {
		case "+", "-", "*", "/":
			top := s2[len(s2)-1]
			for greaterOrEqual(top, val) {
				s2 = s2[:len(s2)-1]
				s1 = append(s1, top)
				top = s2[len(s2)-1]
			}
			s2 = append(s2, val)
		default:
			s1 = append(s1, val)
		}
	}
	for i := len(s2) - 1; i > 0; i-- {
		s1 = append(s1, s2[i])
	}
	return s1
}

func greaterOrEqual(operator1, operator2 string) bool {
	return pririoty(operator1) >= pririoty(operator2)
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
