package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	lenth := 0
	tmp := x
	for ; tmp > 0; tmp = tmp / 10 {
		lenth++
	}
	forwardx, backwardx := x, x
	fx := forwardx
	fxt := 1
	for i := 0; i < lenth/2; i = i + 1 {
		fx = forwardx
		fxt = 1
		for j := i + 1; j < lenth; j++ {
			fx = fx / 10
			fxt = fxt * 10
		}
		forwardx = forwardx % fxt
		bx := backwardx % 10
		if fx != bx {
			return false
		}
		backwardx = backwardx / 10
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome2(1241421))
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	lenth := len(str)
	for i := 0; i < lenth/2; i++ {
		if str[i] != str[lenth-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	lenth := 0
	tmp := x
	revertedNumber := 0
	for ; tmp > 0; tmp = tmp / 10 {
		revertedNumber = revertedNumber*10 + tmp%10
		lenth++
	}
	return revertedNumber == x
}
