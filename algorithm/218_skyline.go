package algorithm

//GetSkyline ...
func GetSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}

	res := make([][]int, 2*len(buildings))
	for i := 0; i < len(buildings); i++ {
		res[2*i] = []int{buildings[i][0], buildings[i][2], 0}
		res[2*i+1] = []int{buildings[i][1], buildings[i][2], 1}
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res)-i-1; j++ {
			if res[j][0] > res[j+1][0] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}

	left := make([][]int, 0)
	valid := make([][]int, 0)
	cur := res[0]
	left = append(left, cur)
	for i := 1; i < len(res); i++ {
		next := res[i]
		if cur[2] == 0 {
			//左坐标点
			if next[2] == 0 {
				left = append(left, next)
				if cur[1] <= next[1] {
					valid = append(valid, cur)
					cur = next
				}
			} else {
				if cur[1] == next[1] {
					valid = append(valid, cur)
					cur = next
				}
				subLeft(&left, next[1])
			}
		} else {
			if next[2] == 1 {
				if cur[1] == next[1] {
					cur = next
				} else {
					top := topLeft(&left)
					cur = []int{cur[0], top, 0}
					if top == next[1] {
						valid = append(valid, cur)
						cur = next
					}
				}
				subLeft(&left, next[1])
			} else {
				left = append(left, next)
				valid = append(valid, []int{cur[0], 0, 0})
				cur = next
			}
		}
	}
	//最后一点
	valid = append(valid, []int{res[len(res)-1][0], 0, 0})

	valid2 := make([][]int, 0)
	for i := 0; i < len(valid)-1; i++ {
		if valid[i][0] == valid[i+1][0] {
			//同横坐标
			if valid[i+1][1] < valid[i][1] {
				valid[i+1][1] = valid[i][1]
			}
			continue
		}
		valid2 = append(valid2, valid[i][0:2])
	}
	//最后一个
	valid2 = append(valid2, valid[len(valid)-1][0:2])

	valid = valid2
	valid2 = valid2[:0]
	for i := 0; i < len(valid)-1; i++ {
		if valid[i][1] == valid[i+1][1] {
			//同高
			valid[i+1][0] = valid[i][0]
			continue
		}
		valid2 = append(valid2, valid[i][0:2])
	}
	v := valid[len(valid)-1][0:2]
	v[1] = 0
	valid2 = append(valid2, v)

	return valid2
}

func subLeft(left *[][]int, h int) {
	for i := 0; i < len(*left); i++ {
		if (*left)[i][1] == h {
			if i+1 == len(*left) {
				(*left) = (*left)[0:i]
			} else {
				(*left) = append((*left)[0:i], (*left)[i+1:]...)
			}
			return
		}
	}
}

func topLeft(left *[][]int) int {
	max := 0
	for _, l := range *left {
		if l[1] > max {
			max = l[1]
		}
	}
	return max
}
