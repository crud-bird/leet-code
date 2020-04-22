package algorithm

import (
	"math"
)

//29.两数相除
func divide(dividend int, divisor int) int {
	//int32范围，-2^31~2^31-1
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	a := dividend
	b := divisor
	if dividend < 0 {
		a = -a
	}
	if divisor < 0 {
		b = -b
	}

	res := 0
	for a >= b {
		tmp := b
		cnt := 1
		for a >= tmp<<1 {
			tmp = tmp << 1
			cnt = cnt << 1
		}
		a -= tmp
		res += cnt
	}

	if (divisor > 0 && dividend > 0) || (dividend < 0 && divisor < 0) {
		return res
	}

	return -res
}
