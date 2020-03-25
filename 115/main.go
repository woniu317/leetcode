package main

import (
	"fmt"
)

func main() {
	s := "rabbbit"
	t := "rabbit"
	s = "babgbag"
	t = "bag"
	//s = "adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc"
	//t = "bcddceeeebecbc"
	fmt.Println(numDistinct(s, t))
}

func numDistinct(s string, t string) int {
	//var num int = 0
	//findSeq(s, t, &num)
	//return num
	return dpFindSeq(s, t)
}

//递归化
func findSeq(s, t string, num *int) {
	if len(t) == 0 {
		*num++
		return
	}
	var j int = 0
	for i := 0; i < len(s); i++ {
		if s[i] == t[j] {
			findSeq(s[i+1:], t[j+1:], num)
			continue
		}
	}
}

func dpFindSeq(s, t string) int {
	dp := make([][]int, len(t)+1)
	for i := 0; i < len(t)+1; i++ {
		dp[i] = make([]int, len(s)+1)
	}
	for i := 0; i < len(s)+1; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < len(t)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}
