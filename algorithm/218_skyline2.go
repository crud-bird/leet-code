package algorithm

//GetSkyline ...
func GetSkyline2(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}

	points := make([][]int, 2*len(buildings))
	for i := 0; i < len(buildings); i++ {
		points[2*i] = []int{buildings[i][0], buildings[i][2], 0}
		points[2*i+1] = []int{buildings[i][1], buildings[i][2], 1}
	}
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points)-i-1; j++ {
			if points[j][0] > points[j+1][0] {
				points[j], points[j+1] = points[j+1], points[j]
			}
		}
	}

	h := Heap{-1, 0}
	res := [][]int{}
	for i := 0; i < len(points); i++ {
		if points[i][2] == 0 {
			if points[i][1] > h.Top() {
				res = append(res, []int{points[i][0], points[i][1]})
			}
			// height = add(&height, points[i][1])
			h.Add(points[i][1])
		} else {
			t := h.Top()
			// sub(&height, points[i][1])
			h.Remove(points[i][1])
			if t != h.Top() {
				res = append(res, []int{points[i][0], h.Top()})
			}
		}
	}

	//合并横坐标相同的点
	valid := [][]int{res[0]}
	for i, j := 1, 0; i < len(res); i++ {
		if valid[j][0] == res[i][0] {
			if (i == len(res)-1 && valid[j][1] > res[i][1]) || (i < len(res)-1 && valid[j][1] < res[i][1]) {
				valid[j] = res[i]
			}
			continue
		}
		valid = append(valid, res[i])
		j++
	}

	//合并相邻的相同高度的点
	res = valid
	valid = [][]int{res[0]}
	for i, j := 1, 0; i < len(res); i++ {
		if valid[j][1] == res[i][1] {
			continue
		}
		valid = append(valid, res[i])
		j++
	}

	return valid
}

type Heap []int

func (h *Heap) Top() int {
	if len(*h) > 1 {
		return (*h)[1]
	}
	return -1
}

func (h *Heap) Add(e int) {
	*h = append(*h, e)
	r := len(*h) - 1
	h.up(r)
}

func (h *Heap) Remove(e int) {
	r := h.search(1, e)
	if r > 0 {
		(*h)[r] = (*h)[len(*h)-1]
		*h = (*h)[:len(*h)-1]
		if r < len(*h) {
			h.up(r)
			h.down(r)
		}
	}
}

func (h *Heap) search(r int, e int) int {
	if r < len(*h) && (*h)[r] == e {
		return r
	}

	next := r << 1
	if next < len(*h) && (*h)[next] >= e {
		if res := h.search(next, e); res > 0 {
			return res
		}
	}
	next++
	if next < len(*h) && (*h)[next] >= e {
		if res := h.search(next, e); res > 0 {
			return res
		}
	}

	return -1
}

func (h *Heap) down(r int) {
	for r < len(*h) {
		next := r << 1
		if next+1 < len(*h) && (*h)[next+1] > (*h)[next] && (*h)[next+1] > (*h)[r] {
			(*h)[r], (*h)[next+1] = (*h)[next+1], (*h)[r]
			r = next + 1
			continue
		}
		if next < len(*h) && (*h)[next] > (*h)[r] {
			(*h)[r], (*h)[next] = (*h)[next], (*h)[r]
			r = next
			continue
		}

		return
	}
}

func (h *Heap) up(r int) {
	for r > 1 {
		pre := r >> 1
		if (*h)[pre] >= (*h)[r] {
			return
		}
		(*h)[r], (*h)[pre] = (*h)[pre], (*h)[r]
		r = pre
	}
}
