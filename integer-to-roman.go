package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}

type ScaleRoman struct {
	scale int
	roman string
}

var scaleList = []ScaleRoman{
	{1000, "M"},
	{500, "D"},
	{100, "C"},
	{50, "L"},
	{10, "X"},
	{5, "V"},
	{1, "I"},
}

var romanMap = map[int]string{
	4:   "IV",
	9:   "IX",
	40:  "XL",
	90:  "XC",
	400: "CD",
	900: "CM",
}

func intToRoman(num int) string {
	var result = ""
	for index, value := range scaleList {
		//先校验是否是用减法
		if index%2 == 1 && num*10/scaleList[index-1].scale == 9 {
			result += romanMap[9*scaleList[index-1].scale/10]
			num -= 9 * scaleList[index-1].scale / 10
		}
		remainder := num / value.scale
		num = num % value.scale
		if remainder == 0 {
			continue
		}
		if va, ok := romanMap[remainder*value.scale]; ok {
			result += va
		} else {
			for i := 0; i < remainder; i++ {
				result += value.roman
			}
		}
	}
	return result
}
