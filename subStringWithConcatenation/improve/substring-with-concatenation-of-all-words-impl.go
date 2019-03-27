package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if words == nil || len(words) == 0 {
		return nil
	}
	var result = make([]int, 0)
	wordsMap := sliceToMap(words)
	wordLen := len(words[0])
	slen := len(s)

	temMap := make(map[string]int)
	for i := 0; i < wordLen; i++ {
		copyMap(wordsMap, temMap)
		left := i
		matchWord := len(words)
		for j := i; j <= slen-wordLen; j += wordLen {
			word := s[j : j+wordLen]
			if value, ok := temMap[word]; ok {
				temMap[word] = value - 1
				matchWord--
				for temMap[word] < 0 {
					str1 := s[left : left+wordLen]
					temMap[str1]++
					matchWord++
					left += wordLen
				}
				if matchWord == 0 {
					result = append(result, left)
					temMap[s[left:left+wordLen]] += 1
					matchWord++
					left += wordLen
				}
			} else {
				left = j + wordLen
				matchWord = len(words)
				copyMap(wordsMap, temMap)
			}

		}
	}
	return result
}

func sliceToMap(words []string) map[string]int {
	var wordsMap = make(map[string]int)
	for _, word := range words {
		wordsMap[word] = wordsMap[word] + 1
	}
	return wordsMap
}

func copyMap(src, target map[string]int) {
	for key, value := range src {
		target[key] = value
	}
}

//"wordgoodgoodgoodbestword"
//["word","good","best","good"]

//"foobarfoobar"
//["foo","bar"]

func main() {
	//fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	//fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	fmt.Println(findSubstring("foobarfoobar", []string{"foo", "bar"}))
}
