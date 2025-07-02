package arrays

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	
	l := len(strs[0])

	for i := range len(strs) - 1 {
		l = min(l, len(strs[i+1]))
		for j := range l {
			if strs[i][j] != strs[i+1][j] {
				l = min(len(strs[i][:j]), l)
				continue
			}
		}
		if l == 0 {
			break
		}
	}
	return strs[0][:l]
}
