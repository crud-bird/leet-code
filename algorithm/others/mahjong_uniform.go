package others

//MahjongUniform 麻将清一色，输入数据已排序
func MahjongUniform(arr []int) string {
	//todo 3个连续的对子可以作为两个顺子
	l := len(arr)
	if l != 2 && l != 5 && l != 8 && l != 11 && l != 14 {
		return "no"
	}

	pairs := make([]int, 0)
	//找对子
	pre, cur := 0, 0
	for i, e := range arr {
		pre, cur = cur, e
		if pre == cur {
			pairs = append(pairs, i)
		}
	}

	for _, pairidx := range pairs {
		tmp := make([]int, 0)
		for i, e := range arr {
			if i == pairidx || i == pairidx-1 {
				continue
			}
			tmp = append(tmp, e)
		}

		//初始值-10，避免和实际的牌产生联系
		var pre, cur, next int = -10, -10, -10
		for i := 0; i < len(tmp); i++ {
			pre, cur, next = cur, next, tmp[i]
			if (pre == cur && cur == next) || (pre == cur-1 && cur == next-1) {
				//有用的牌标记为100
				tmp[i-2], tmp[i-1], tmp[i] = 100, 100, 100
				pre, cur, next = -10, -10, -10
				continue
			}
		}

		res := true
		for _, e := range tmp {
			if e != 100 {
				res = false
				break
			}
		}
		if res {
			return "yes"
		}

	}

	return "no"
}
