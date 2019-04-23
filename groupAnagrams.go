package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string)

	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		sortStr := strings.Join(chars, "")
		resultMap[sortStr] = append(resultMap[sortStr], str)
	}

	var result [][]string
	for _, re := range resultMap {
		result = append(result, re)
	}
	return result
}
