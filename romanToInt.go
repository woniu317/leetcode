package main

import "fmt"

//性能romanToInt=romanToInt2=romanToInt1
func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

var romanScaleMap = map[string]int{
	"M":  1000,
	"D":  500,
	"C":  100,
	"L":  50,
	"X":  10,
	"V":  5,
	"I":  1,
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt1(s string) int {
	result := 0
	lenth := len(s)
	for i := 0; i < lenth; {
		if i != lenth-1 {
			if value := romanScaleMap[s[i:i+2]]; value != 0 {
				result += value
				i += 2
				continue
			}
		}
		result += romanScaleMap[s[i:i+1]]
		i++
	}
	return result
}

func romanToInt(s string) int {
	result := 0
	lenth := len(s)
	for i := 0; i < lenth; {
		rightEdge := min2(i+2, lenth)
		value, step := getScaleValue(s[i:rightEdge])
		result += value
		i += step
	}
	return result
}

func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func romanToInt2(s string) int {
	result := 0
	lenth := len(s)
	for i := 0; i < lenth; {
		rightEdge := min2(i+2, lenth)
		switch s[i:rightEdge] {
		case "IV":
			result += 4
			i += 2
		case "IX":
			result += 9
			i += 2
		case "XL":
			result += 40
			i += 2
		case "XC":
			result += 90
			i += 2
		case "CD":
			result += 400
			i += 2
		case "CM":
			result += 900
			i += 2
		default:
			switch s[i : i+1] {
			case "M":
				result += 1000
				i++
			case "D":
				result += 500
				i++
			case "C":
				result += 100
				i++
			case "L":
				result += 50
				i++
			case "X":
				result += 10
				i++
			case "V":
				result += 5
				i++
			case "I":
				result += 1
				i++
			}
		}
	}
	return result
}

func getScaleValue(key string) (value, step int) {
	switch key {
	case "IV":
		return 4, 2
	case "IX":
		return 9, 2
	case "XL":
		return 40, 2
	case "XC":
		return 90, 2
	case "CD":
		return 400, 2
	case "CM":
		return 900, 2
	default:
		switch key[0:1] {
		case "M":
			return 1000, 1
		case "D":
			return 500, 1
		case "C":
			return 100, 1
		case "L":
			return 50, 1
		case "X":
			return 10, 1
		case "V":
			return 5, 1
		case "I":
			return 1, 1
		}
	}
	return 0, 1
}
