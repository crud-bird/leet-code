package others_test

import (
	"fmt"
	"testing"

	"github.com/crud-bird/leet-code/algorithm/others"
)

func TestDicing(t *testing.T) {
	dice := others.Dicing("RA")
	fmt.Printf("%v", dice[1:])
}
