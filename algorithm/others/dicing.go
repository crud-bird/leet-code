package others

//Dicing 掷骰子
func Dicing(operation string) []int {
	dice := []int{0, 1, 2, 3, 4, 5, 6} //左，右，前，后，上，下
	for i := 0; i < len(operation); i++ {
		dicing(dice, operation[i])
	}
	return dice
}

func dicing(dice []int, op byte) {
	switch op {
	case 'L': //左
		dice[1], dice[2], dice[5], dice[6] = dice[5], dice[6], dice[2], dice[1]
	case 'R': //右
		dice[1], dice[2], dice[5], dice[6] = dice[6], dice[5], dice[1], dice[2]
	case 'F': //前
		dice[3], dice[4], dice[5], dice[6] = dice[5], dice[6], dice[4], dice[3]
	case 'B': //后
		dice[3], dice[4], dice[5], dice[6] = dice[6], dice[5], dice[3], dice[4]
	case 'A': //逆时针
		dice[1], dice[2], dice[3], dice[4] = dice[4], dice[3], dice[1], dice[2]
	case 'C': //顺时针
		dice[1], dice[2], dice[3], dice[4] = dice[3], dice[4], dice[2], dice[1]
	}
}
