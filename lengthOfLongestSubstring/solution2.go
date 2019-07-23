package lengthOfLongestSubstring

func Solution2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	subMax := Solution2(s[1:])
	sub := make(map[byte]byte)
	length := 0
	for i, _ := range s {
		if _, ok := sub[s[i]]; ok {
			break
		}
		sub[s[i]] = s[i]
		length++
	}
	if length > subMax {
		return length
	}
	return subMax
}
