package others_test

import (
	"testing"

	"github.com/crud-bird/leet-code/algorithm/others"
)

func TestLISP(t *testing.T) {
	str := "(search (combine (reverse ymk) (quote aa) ko) ya)"
	println(others.LISP(str))
}
