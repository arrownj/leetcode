package main

func RomanToInt(roman string) int {
	adding := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	subtract := map[byte][]byte{
		'I': []byte{'V', 'X'},
		'X': []byte{'L', 'C'},
		'C': []byte{'D', 'M'},
	}
	value := 0
	for i := 0; i < len(roman); i++ {
		c := roman[i]
		value += adding[c]
		if val, ok := subtract[c]; ok {
			if i < len(roman)-1 && contains(val, roman[i+1]) {
				value -= adding[c]
				value -= adding[c]
			}
		}

	}
	return value
}

func contains(slice []byte, value byte) bool {
	for _, c := range slice {
		if c == value {
			return true
		}
	}
	return false
}
