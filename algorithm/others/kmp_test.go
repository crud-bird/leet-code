package others_test

import (
	"testing"

	"github.com/crud-bird/leet-code/algorithm/others"
)

func TestKMP(t *testing.T) {
	s := "ssfesifdjaskd"
	p := "askd"
	println(others.KMP(s, p))

	s = "tyuiohdskk"
	p = "skk"
	println(others.KMP(s, p))
}
