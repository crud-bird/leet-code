package others

import "fmt"

//SortWords 单词排序
func SortWords(words []string) {
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words)-i-1; j++ {
			if wordsLess(words[j], words[j+1], 0) {
				continue
			}
			words[j], words[j+1] = words[j+1], words[j]
		}
	}
	for _, word := range words {
		fmt.Printf(" %s", word)
	}
}

func wordsLess(str1, str2 string, idx int) bool {
	if idx == len(str1) || idx == len(str2) {
		return len(str1) < len(str2)
	}

	if str1[idx] == str2[idx] {
		return wordsLess(str1, str2, idx+1)
	}

	return str1[idx] < str2[idx]
}
