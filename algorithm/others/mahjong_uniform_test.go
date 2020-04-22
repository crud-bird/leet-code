package others_test

import (
	"testing"

	"github.com/crud-bird/leet-code/algorithm/others"
)

func TestMahjongUniform(t *testing.T) {
	// arr := []int{2, 2, 2, 3, 4, 5, 5, 5, 5, 6, 7, 8, 9, 9}
	// arr := []int{5, 5, 5, 6, 7, 8, 9, 9}
	arr := []int{5, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	println(others.MahjongUniform(arr))
}
