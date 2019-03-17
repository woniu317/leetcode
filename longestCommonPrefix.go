package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = longestCommonPrefixTwoString(strs[i], prefix)
	}
	return prefix
}

func longestCommonPrefixTwoString(str1, str2 string) string {
	if len(str1) == 0 || len(str2) == 0 {
		return ""
	}
	for i, j := 0, 0; i < len(str1) && j < len(str2); {
		if str1[i] != str2[j] {
			return str1[0:i]
		}
		i += 1
		j += 1
	}
	if len(str1) < len(str2) {
		return str1
	} else {
		return str2
	}
}
