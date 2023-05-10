package longest_common_prefix

func GetLongestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	i := 0
	for ; i < len(strs[0]); i++ {
		prefix := strs[0][:i+1]
		break_out := false
		for _, str := range strs {
			if i+1 > len(str) || str[:i+1] != prefix {
				break_out = true
				break
			}
		}
		if break_out {
			break
		}
	}
	return strs[0][:i]
}
