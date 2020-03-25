package main

import (
	"fmt"
	"math"
	"strings"
)

func reverseWords(s string) string {
	lastSpaceIndex := len(s)
	var res strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if i == lastSpaceIndex-1 {
				lastSpaceIndex--
				continue
			}
			res.WriteString(s[i+1 : lastSpaceIndex])
			res.WriteString(" ")
			lastSpaceIndex = i
		} else if i == 0 {
			res.WriteString(s[i:lastSpaceIndex])
			res.WriteString(" ")
		}
	}
	return strings.TrimRight(res.String(), " ")
}

func main() {
	fmt.Println(math.MinInt64)
	fmt.Println("==" + reverseWords("the sky is blue") + "==")
}
func maxProduct(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {

		tmp := 1
		for k := i; k <= len(nums); k++ {
			tmp *= nums[k]
			if tmp > max {
				max = tmp
			}
		}

	}
	return max
}
