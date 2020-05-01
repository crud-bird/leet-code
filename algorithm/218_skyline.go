package algorithm

func GetSkyline(buildings [][]int) [][]int {
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

	height := []int{0}
	res := [][]int{}
	for i := 0; i < len(points); i++ {
		if points[i][2] == 0 {
			if points[i][1] > top(&height) {
				res = append(res, []int{points[i][0], points[i][1]})
			}
			height = add(&height, points[i][1])
		} else {
			t := top(&height)
			sub(&height, points[i][1])
			if t != top(&height) {
				res = append(res, []int{points[i][0], top(&height)})
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

func sub(s *[]int, h int) {
	for i := len(*s) - 1; i >= 0; i-- {
		if (*s)[i] == h {
			if i == len(*s)-1 {
				*s = (*s)[:i]
			} else {
				*s = append((*s)[:i], (*s)[i+1:]...)
			}
			return
		}
	}
}

func top(s *[]int) int {
	return (*s)[len(*s)-1]
}

func add(s *[]int, e int) []int {
	var i int
	for i = 0; i < len(*s); i++ {
		if e < (*s)[i] {
			break
		}
	}
	if i == len(*s) {
		*s = append(*s, e)
		return *s
	}
	res := make([]int, 0, len(*s)+1)
	res = append(res, (*s)[:i]...)
	res = append(res, e)
	res = append(res, (*s)[i:]...)
	return res
}
