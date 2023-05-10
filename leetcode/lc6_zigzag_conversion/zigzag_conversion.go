package zigzag_conversion

func ConvertZigzag(s string, num int) string {
	if num <= 1 {
		return s
	}

	str := []rune(s)
	ret := []rune{}
	for i := 0; i < num; i++ {
		for index := i; index < len(str); {
			ret = append(ret, str[index])
			if i != 0 && i != num-1 {
				next := index + 2*(num-1) - 2*i
				if next < len(str) {
					ret = append(ret, str[next])
				}
			}
			index += 2 * (num - 1)
		}
	}
	return string(ret)
}
