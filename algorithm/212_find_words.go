package algorithm

import (
	"fmt"
)

func findWords(board [][]byte, words []string) []string {
	res := make([]string, 0)
	sortWords(words)
	for _, word := range words {
		for i := 0; i < len(board); i++ {
			found := false
			for j := 0; j < len(board[i]); j++ {
				if match(board, i, j, 0, word, make(map[string]bool)) {
					res = append(res, word)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	return res
}

func match(board [][]byte, i, j, cur int, word string, blocked map[string]bool) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
		return false
	}
	if board[i][j] != byte(word[cur]) {
		return false
	}

	key := fmt.Sprintf("%d+%d", i, j)
	blocked[key] = true
	cur++
	if cur == len(word) {
		return true
	}

	key = fmt.Sprintf("%d+%d", i+1, j)
	if !blocked[key] {
		if match(board, i+1, j, cur, word, blocked) {
			return true
		}
	}

	key = fmt.Sprintf("%d+%d", i-1, j)
	if !blocked[key] {
		if match(board, i-1, j, cur, word, blocked) {
			return true
		}
	}

	key = fmt.Sprintf("%d+%d", i, j-1)
	if !blocked[key] {
		if match(board, i, j-1, cur, word, blocked) {
			return true
		}
	}

	key = fmt.Sprintf("%d+%d", i, j+1)
	if !blocked[key] {
		if match(board, i, j+1, cur, word, blocked) {
			return true
		}
	}

	blocked[fmt.Sprintf("%d+%d", i, j)] = false
	return false
}

func sortWords(words []string) {
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words)-i-1; j++ {
			if wordsLess(words[j], words[j+1], 0) {
				continue
			}
			words[j], words[j+1] = words[j+1], words[j]
		}
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
