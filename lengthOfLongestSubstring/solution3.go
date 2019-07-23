package lengthOfLongestSubstring

func Solution3(s string) int {
	max, i, j := 0, 0, 0
	sub := make(map[byte]int)
	for i < len(s) && j < len(s) {
		if _, ok := sub[s[j]]; ok {
			delete(sub, s[i])
			i++
		} else {
			sub[s[j]] = j
			j++
			if j-i > max {
				max = j - i
			}
		}
	}
	return max
}
