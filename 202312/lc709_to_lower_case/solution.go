package main

func toLowerCase(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			result[i] = byte(int(s[i]) - int('A'-'a'))
		} else {
			result[i] = s[i]
		}
	}
	return string(result)
}
