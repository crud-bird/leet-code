package others

import "fmt"

//SimpleArithmetic ...
// 问题描述：
// 输入一个只包含个位数字的简单四则运算表达式字符串，计算该表达式的值
// 注：
// 1、表达式只含 +, -, *, / 四则运算符，不含括号
// 2、表达式数值只包含个位整数(0-9)，且不会出现0作为除数的情况
// 3、要考虑加减乘除按通常四则运算规定的计算优先级
// 4、除法用整数除法，即仅保留除法运算结果的整数部分。比如8/3=2。输入表达式保证无0作为除数情况发生
// 5、输入字符串一定是符合题意合法的表达式，其中只包括数字字符和四则运算符字符，除此之外不含其它任何字符，不会出现计算溢出情况
// ————————————————
// 原文链接：https://blog.csdn.net/wangjian8855/java/article/details/9851557
func SimpleArithmetic(str string) int {
	var left, right = int(str[0] - '0'), 0
	var op byte
	for i := 1; i < len(str); i = i + 2 {
		op = str[i]
		right = int(str[i+1] - '0')
		if len(str) == i+2 {
			return arithmetic(left, right, op)
		}

		if op == '*' || op == '/' || str[i+2] == '+' || str[i+2] == '-' {
			left = arithmetic(left, right, op)
			continue
		}

		return arithmetic(left, SimpleArithmetic(str[i+1:]), op)
	}

	return 0
}

func arithmetic(a, b int, op byte) int {
	fmt.Printf("%d %s %d\n", a, string(op), b)
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}
