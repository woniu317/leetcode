package main

import (
	"fmt"
	"strconv"
)

func explanExpress(str string) []string {
	res := make([]string, 0)
	for i := 0; i < len(str); {
		if str[i] >= '0' && str[i] <= '9' {
			j := i
			for ; j < len(str); j++ {
				if str[j] < '0' || str[j] >= '9' {
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
	fmt.Println(expValue("1+2*3/2+1-10"))
}

func expValue(str string) int {
	express := explanExpress(str)
	ds := make([]int, 0)
	os := make([]string, 0)
	for i := 0; i < len(express); i++ {
		if express[i][0] >= '0' && express[i][0] <= '9' {
			tmp, _ := strconv.Atoi(express[i])
			ds = append(ds, tmp)
		} else {
			if len(os) == 0 || great(express[i], os[len(os)-1]) {
				os = append(os, express[i])
			} else {
				prossess(&ds, &os, express[i])
				os = append(os, express[i])
			}
		}
	}
	prossess(&ds, &os, "#")
	return ds[0]
}

func prossess(ds *[]int, os *[]string, newOperator string) {
	for len(*os) != 0 && !great(newOperator, (*os)[len(*os)-1]) {
		n := len(*ds)
		d1 := (*ds)[n-1]
		d2 := (*ds)[n-2]
		operator := (*os)[len(*os)-1]
		res := cal(d2, d1, operator)
		(*ds)[n-2] = res
		*ds = (*ds)[:n-1]
		*os = (*os)[:len(*os)-1]
	}
}

func cal(d1 int, d2 int, operator string) int {
	switch operator {
	case "+":
		return d1 + d2
	case "-":
		return d1 - d2
	case "*":
		return d1 * d2
	default:
		return d1 / d2
	}
}

func great(operator1 string, operator2 string) bool {
	return priority(operator1) > priority(operator2)
}

func priority(oper1 string) int {
	switch oper1 {
	case "#":
		return 0
	case "+", "-":
		return 1
	default:
		return 2
	}
}
