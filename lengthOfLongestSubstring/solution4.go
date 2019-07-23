package lengthOfLongestSubstring

func Solution4(s string) int {
	max, i, j := 0, 0, 0
	sub := make(map[byte]int)
	for i < len(s) && j < len(s) {
		if idx, ok := sub[s[j]]; ok {
			if idx+1 > i {
				i = idx + 1
			}
		}
		sub[s[j]] = j
		j++
		if j-i > max {
			max = j - i
		}
	}
	return max
}
