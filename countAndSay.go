package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(4))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	count := int64(0)
	cur := str[0]
	result := []byte{}
	for i := 0; i < len(str); i++ {
		if cur == str[i] {
			count++
		} else {
			result = strconv.AppendInt(result, count, 10)
			result = append(result, cur)
			count = 1
			cur = str[i]
		}
	}
	if count > 0 {
		result = strconv.AppendInt(result, count, 10)
		result = append(result, cur)
	}
	return string(result)
}
