package lc1768_merge_alternately

func Solution(s1, s2 string) string {
	merged := make([]byte, len(s1)+len(s2), len(s1)+len(s2))
	minLength := len(s1)
	if len(s1) > len(s2) {
		minLength = len(s2)
	}
	i := 0
	for ; i < minLength; i++ {
		merged[i*2] = s1[i]
		merged[i*2+1] = s2[i]
	}
	for j := minLength; j < len(s1); j++ {
		merged[minLength+j] = s1[j]
	}
	for j := minLength; j < len(s2); j++ {
		merged[minLength+j] = s2[j]
	}
	return string(merged)
}
