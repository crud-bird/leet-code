package others_test

import (
	"testing"

	"github.com/crud-bird/leet-code/algorithm/others"
)

func TestSortWords(t *testing.T) {
	strs := []string{"hello", "world", "enjoy", "enable", "end", "words", "hehe", "word"}
	others.SortWords(strs)
}
