package others

import (
	"strings"
)

// 题目描述
// LISP语言唯一的语法就是括号要配对。
// 形如 (OP P1 P2 …)，括号内元素由单个空格分割。
// 其中第一个元素OP为操作符，后续元素均为其参数，参数个数取决于操作符类型
// 注意：参数 P1, P2 也有可能是另外一个嵌套的 (OP P1 P2 …)
// 当前OP类型为 quote / reverse / search / combine 字符串相关的操作：
// - quote: 引用一个字符串，即返回字符串本身内容
// 参数个数 1
// - reverse: 把字符串反转，并返回
// 参数个数 1
// - search: 在第一个字符串中查找第二个字符串的第一次出现，返回从这开始到结束的所有字符串
// 如果查找不到，返回空字符串
// 参数个数 2
// - combine: 把所有字符串组合起来
// 参数个数不定，但至少 1 个
// 其中P1, P2 等参数可能是带双引号的字符串,如 “abc”，也有可能是另外一个 (OP P1 P2 …)
// 上述字符串包括引号；引号中间的所有字符，均为 ASCII 可打印字符，且不会再出现引号 (“)
// 输出也为带双引号的字符串
// ————————————————
// 原文链接：https://blog.csdn.net/HappyKocola/java/article/details/73864891

//LISP 简化处理，输入参数不包含空格和括号,参数不用引号
func LISP(str string) string {
	strs := strings.Split(str, " ")
	stack := make([]string, 0)
	for _, s := range strs {
		if idx := strings.Index(s, ")"); idx > -1 {
			var i int
			for i = len(stack) - 1; i >= 0; i-- {
				if ii := strings.Index(stack[i], "("); ii > -1 {
					break
				}
			}
			args := stack[i+1:]
			op := strings.Replace(stack[i], "(", "", 1)
			stack = stack[:i]
			args = append(args, strings.Replace(s, ")", "", 1))
			stack = append(stack, operate(op, args...))
			continue
		}

		//入栈
		stack = append(stack, s)
	}

	return stack[0]
}

func operate(op string, args ...string) string {
	switch op {
	case "reverse":
		chars := []byte(args[0])
		for i, k := 0, len(chars)-1; i < k; i, k = i+1, k-1 {
			chars[i], chars[k] = chars[k], chars[i]
		}
		return string(chars)
	case "quote":
		return args[0]
	case "combine":
		return strings.Join(args, "")
	case "search":
		idx := strings.Index(args[0], args[1])
		if idx == -1 {
			return ""
		}
		return string([]byte(args[0])[idx:])
	}

	return ""
}
