package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

var letterMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return make([]string, 0)
	}
	result := make([]string, 1)
	digitsLen := len(digits)
	for i := 0; i < digitsLen; i++ {
		digit := digits[i : i+1]
		targetStrLen := len(letterMap[digit])
		result = timesArray(result, targetStrLen, letterMap[digit])
	}
	return result
}

func timesArray(source []string, times int, appendStr string) []string {
	sourceLen := len(source)
	targetLen := sourceLen * times
	targetArray := make([]string, targetLen)
	for i := 0; i < times; i++ {
		for j := 0; j < sourceLen; j++ {
			targetArray[i*sourceLen+j] = source[j] + appendStr[i:i+1]
		}
	}
	return targetArray
}
