package algorithm

//678.有效的括号字符串
func checkValidString(s string) bool {
	return doCheckValidString(s, 0)
}

func doCheckValidString(s string, pairs int) bool {
	for i := 0; i < len(s) && pairs >= 0; i++ {
		if s[i] == '(' {
			pairs++
			continue
		} else if s[i] == ')' {
			pairs--
			continue
		}

		if i == len(s) {
			break
		}

		sub := s[i+1:]
		return doCheckValidString(sub, pairs) || doCheckValidString(sub, pairs+1) || doCheckValidString(sub, pairs-1)
	}

	return pairs == 0
}
