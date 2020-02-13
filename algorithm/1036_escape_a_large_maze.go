package algorithm

import (
	"fmt"
)

//lettcode第1036题，逃离大迷宫
func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	//只检查sourece到target会存在target被包围的情况
	return walk(blocked, source, target) && walk(blocked, target, source)
}

func walk(blocked [][]int, source []int, target []int) bool {
	if len(blocked) < 2 {
		return true
	}

	//blocked能包围的最大面积，以边界为边的直角三角形
	max := len(blocked) * len(blocked) / 2

	//转成map提高查找效率
	bMap := make(map[string]bool)
	for _, arr := range blocked {
		bMap[fmt.Sprintf("%d,%d", arr[0], arr[1])] = true
	}

	//走过的路径
	route := [][]int{{source[0], source[1]}}
	for len(route) > 0 {
		if source[0] == target[0] && source[1] == target[1] {
			return true
		}

		//走过的面积大于能封锁的最大面积，说明source点已经走出包围
		if len(route) > max {
			return true
		}

		//走过的点设为障碍点，防止回退陷入死循环
		bMap[fmt.Sprintf("%d,%d", source[0], source[1])] = true

		//向右走
		if source[0]+1 < 1000000 && !bMap[fmt.Sprintf("%d,%d", source[0]+1, source[1])] {
			route = append(route, []int{source[0], source[1]})
			source[0]++
			continue
		}

		//向下走
		if source[1]+1 < 1000000 && !bMap[fmt.Sprintf("%d,%d", source[0], source[1]+1)] {
			route = append(route, []int{source[0], source[1]})
			source[1]++
			continue
		}

		//向左走
		if source[0]-1 >= 0 && !bMap[fmt.Sprintf("%d,%d", source[0]-1, source[1])] {
			route = append(route, []int{source[0], source[1]})
			source[0]--
			continue
		}

		//向上走
		if source[1]-1 >= 0 && !bMap[fmt.Sprintf("%d,%d", source[0], source[1]-1)] {
			route = append(route, []int{source[0], source[1]})
			source[1]--
			continue
		}

		//所有方向都走不通，回退到上一步走其他方向，用切片简单实现栈功能
		l := len(route)
		source = route[l-1]
		route = route[:l-1]
	}

	return false
}
