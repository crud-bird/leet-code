package algorithm_test

import (
	"fmt"
	"testing"

	"github.com/crud-bird/leet-code/algorithm"
)

func TestGetSkyline(t *testing.T) {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}, {25, 26, 8}, {27, 28, 8}}
	res := algorithm.GetSkyline(buildings)
	fmt.Printf("%v", res)
}
