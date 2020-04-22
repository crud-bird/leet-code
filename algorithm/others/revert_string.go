package others

import "fmt"

//RevertString 字符串反转
func RevertString(str string) string {
	chars := []byte(str)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	fmt.Println(string(chars))

	for i, k := 0, 0; i < len(chars); i, k = i+k, 0 {
		for chars[i] == ' ' {
			i++
		}

		for i+k < len(chars) && chars[i+k] != ' ' {
			k++
		}

		if k < 2 {
			continue
		}

		for s := 0; s < k/2; s++ {
			chars[i+s], chars[i+k-s-1] = chars[i+k-s-1], chars[i+s]
		}
	}

	return string(chars)
}
