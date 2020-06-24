package others

func KMP(s, p string) int {
	next := kmpHappyPrefix(p)
	j := 0
	for i := 0; i < len(s); i++ {
		for j > 0 && p[j] != s[i] {
			j = next[j-1]
		}
		if p[j] == s[i] {
			j++
		}
		if j == len(p) {
			return i - j + 1
		}
	}
	return -1
}

func kmpHappyPrefix(p string) []int {
	res := make([]int, len(p))
	for i := 1; i < len(p); i++ {
		j := res[i-1]
		for j > 0 && p[j] != p[i] {
			j = res[j-1]
		}

		if p[j] == p[i] {
			res[i] = j + 1
		}
	}

	return res
}
