package main

import "fmt"

func generateParenthesis(n int) []string {
	str := make([]string, 2*n)
	for i := 0; i < n; i++ {
		str[2*i] = "("
		str[2*i+1] = ")"
	}
	return permission(str, 0, 2*n, 0)
}

func main() {
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}

func permission(str []string, start, end int, num int) []string {
	if end == start+1 {
		result := ""
		for _, va := range str {
			result += va
		}
		return []string{result}
	}
	result := make([]string, 0)
	remaindLeft := 0
	for j := 0; j < start; j++ {
		if str[j] == "(" {
			remaindLeft++
		} else {
			remaindLeft--
		}
	}

	if remaindLeft < 0 {
		return result
	} else if remaindLeft == 0 {
		if str[start] == "(" {
			result = append(result[:], permission(str, start+1, end, num+1)[:]...)
			return result
		}
		for i := start + 1; i < end; i++ {
			if str[i] != str[i-1] {
				str[start], str[i] = str[i], str[start]
				result = append(result[:], permission(str, start+1, end, num+1)[:]...)
				str[start], str[i] = str[i], str[start]
				break
			}
		}
	} else {
		for i := start; i < end; i++ {
			if i == start {
				result = append(result[:], permission(str, start+1, end, num+1)[:]...)
				continue
			}
			if str[i] != str[i-1] {
				str[start], str[i] = str[i], str[start]
				result = append(result[:], permission(str, start+1, end, num+1)[:]...)
				str[start], str[i] = str[i], str[start]
				break
			}
		}
	}

	return result
}
