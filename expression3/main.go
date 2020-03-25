package main

import (
	"fmt"
	"strconv"
)

func explan(str string) []string {
	res := make([]string, 0)
	for i := 0; i < len(str); {
		if str[i] >= '0' && str[i] <= '9' {
			j := i
			for ; j < len(str); j++ {
				if str[j] < '0' || str[j] > '9' {
					break
				}
			}
			res = append(res, str[i:j])
			i = j
		} else {
			res = append(res, str[i:i+1])
			i++
		}
	}
	return res
}

func main() {
	fmt.Println(expValue("1+2"))
}

func expValue(str string) int {
	exp := explan(str)
	ds := make([]int, 0)
	os := make([]string, 0)
	for i := 0; i < len(exp); i++ {
		if exp[i][0] >= '0' && exp[i][0] <= '9' {
			value, _ := strconv.Atoi(exp[i])
			ds = append(ds, value)
		} else {
			if len(os) == 0 || greater(exp[i], os[len(os)-1]) {
				os = append(os, exp[i])
			} else {
				prossess(&ds, &os, exp[i])
				os = append(os, exp[i])
			}
		}
	}
	prossess(&ds, &os, "#")
	return ds[0]
}

func prossess(ds *[]int, os *[]string, newOperator string) {
	for len(*os) != 0 && !greater(newOperator, (*os)[len(*os)-1]) {
		//dn := len(*ds)
		//on := len(*os)
		//d1 := (*ds)[dn-1]
		//d2 := (*ds)[dn-2]
		//operator := (*os)[on-1]
		//res := cal(d2, d1, operator)
		//(*ds)[dn-2] = res
		//*ds = (*ds)[:dn-1]
		//*os = (*os)[:on-1]
		l := len(*ds)
		switch (*os)[len(*os)-1] {
		case "+":
			*ds = append((*ds)[:l-2], (*ds)[l-2]+(*ds)[l-1])
		case "-":
			*ds = append((*ds)[:l-2], (*ds)[l-2]-(*ds)[l-1])
		case "*":
			*ds = append((*ds)[:l-2], (*ds)[l-2]*(*ds)[l-1])
		default:
			*ds = append((*ds)[:l-2], (*ds)[l-2]/(*ds)[l-1])
		}
		*os = (*os)[:len(*os)-1]
	}

}

func greater(op1, op2 string) bool {
	return queryPriority(op1) > queryPriority(op2)
}

func queryPriority(operator string) int {
	switch operator {
	case "#":
		return 0
	case "-", "+":
		return 1
	default:
		return 2
	}
}

func cal(data1, data2 int, operator string) int {
	switch operator {
	case "+":
		return data1 + data2
	case "-":
		return data1 - data2
	case "*":
		return data1 * data2
	default:
		return data1 / data2
	}
}
