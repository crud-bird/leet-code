package algorithm_test

import (
	"fmt"
	"testing"

	"github.com/crud-bird/leet-code/algorithm"
)

func TestFindWords2(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}

	res := algorithm.FindWords2(board, words)
	fmt.Printf("%v\n", res)
}
