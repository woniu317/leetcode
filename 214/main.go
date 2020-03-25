package main

import "fmt"

func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func shortestPalindrome(s string) string {
	result := s
	i := len(s)
	for ; i >= 0; i-- {
		if IsPalindrome(s[:i]) {
			break
		}
	}
	tmp := ""
	for j := len(s) - 1; j >= i; j-- {
		tmp += string(s[j])
	}
	return tmp + result
}

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
}
