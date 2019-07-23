package lengthOfLongestSubstring

func Solution1(s string) int {
	maxLength := 0
	for i := 0; i < len(s); i++ {
		sub := make(map[byte]byte)
		length := 0
		for j := i; j < len(s); j++ {
			if _, ok := sub[s[j]]; ok {
				break
			} else {
				sub[s[j]] = s[j]
			}
			length++
		}
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
