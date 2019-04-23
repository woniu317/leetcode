package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Println(isNumber("0e"))
}

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	i, n := 0, len(s)
	if i < n && (s[i] == '+' || s[i] == '-') {
		i++
	}
	isNumeric := false
	for i < n && isDigit(s[i]) {
		i++
		isNumeric = true
	}
	if i < n && s[i] == '.' {
		i++
		for i < n && isDigit(s[i]) {
			i++
			isNumeric = true
		}
	}
	if isNumeric && i < n && s[i] == 'e' {
		i++
		isNumeric = false
		if i < n && (s[i] == '-' || s[i] == '+') {
			i++
		}
		for i < n && isDigit(s[i]) {
			i++
			isNumeric = true
		}
	}
	return isNumeric && i == n
}

func isDigit(b byte) bool {
	digit, _ := utf8.DecodeRune([]byte(string(b)))
	return unicode.IsDigit(digit)
}
