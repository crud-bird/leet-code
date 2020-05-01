package others

//NumberClassiy 数据分类,a四个字节相加模b得到d,d<c则a属于d类；返回最多的类的数字个数
func NumberClassiy(param []int32) int {
	b := param[0]
	c := param[1]
	m := make(map[int32]int)
	for i := 2; i < len(param); i++ {
		a := param[i]
		d := (a&0xff + (a>>8)&0xff + (a>>16)&0xff + (a>>24)&0xff) % b
		if d < c {
			m[d]++
		}
	}

	max := 0
	for _, n := range m {
		if n > max {
			max = n
		}
	}
	return max
}
