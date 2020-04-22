package others

import "fmt"

// Jam不使用阿拉伯数字计数，而是使用小写英文字母计数。他的计数法中，
// 每个数字的位数都是相同的（使用相同个数的字母），英文字母按原先的顺序，
// 排在前面的字母小于排在它后面的字母。我们把这样的“数字”称为Jam数字。在Jam数字中，
// 每个字母互不相同，而且从左到右是严格递增的。每次，Jam还指定使用字母的范围，例如，
// 从2到10，表示只能使用{b,c,d,e,f,g,h,i,j}这些字母。如果再规定位数为5，那么，
// 紧接在Jam数字“bdfij”之后的数字应该是“bdghi”。
// （如果我们用U、V依次表示Jam数字“bdfij”与“bdghi”，则U<V，且不存在Jam数字P，
// 使U<P<V）。你的任务是：对于从文件读入的一个Jam数字，按顺序输出紧接在后面的5个Jam数字，
// 如果后面没有那么多Jam数字，那么有几个就输出几个。
// ————————————————
// 原文链接：https://blog.csdn.net/qq_40590018/java/article/details/84307482

//JamCount jam计数
func JamCount() {
	for {
		var s, t, n int
		var str string
		k, _ := fmt.Scan(&s)
		if k == 0 {
			break
		}
		fmt.Scan(&t)
		fmt.Scan(&n)
		fmt.Scan(&str)
		chars := []byte(str)
		tChar := 'a' + t - 1
		idx := n - 1
		for i := 0; i < 5; i++ {
			if !jamCount(chars, byte(tChar), idx) {
				break
			}
			println(string(chars))
		}
	}
}

func jamCount(chars []byte, t byte, idx int) bool {
	if idx < 0 {
		return false
	}
	if chars[idx]+1 > t {
		exist := jamCount(chars, t-1, idx-1)
		if !exist {
			return false
		}
		chars[idx] = chars[idx-1] + 1
		return true
	}
	chars[idx]++
	return true
}
