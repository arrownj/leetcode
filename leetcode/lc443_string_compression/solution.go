package lc443_string_compression

import "fmt"

func Solution(chars []byte) int {
	outputIndex := 0
	for i := 0; i < len(chars); {
		chars[outputIndex] = chars[i]
		outputIndex++
		charCount := 0
		j := i
		for ; j < len(chars); j++ {
			if chars[j] == chars[i] {
				charCount++
			} else {
				break
			}
		}
		if charCount > 1 {
			numChars := GetNumChars(charCount)
			for k := 0; k < len(numChars); k++ {
				chars[outputIndex] = numChars[k]
				outputIndex++
			}
		}
		i = j
	}
	return outputIndex
}

func GetNumChars(num int) []byte {
	if num == 0 {
		return []byte{'0'}
	}

	value := num
	count := 0
	for value > 0 {
		value = value / 10
		count++
	}
	fmt.Printf("num: %d, count: %d, value:%d\n", num, count, value)
	chars := make([]byte, count, count)

	value = num
	index := count
	for value > 0 {
		digit := value % 10
		value = value / 10
		fmt.Printf("digit %d char is %s\n", digit, string(GetNumChar(digit)))
		chars[index-1] = GetNumChar(digit)
		index--
	}
	fmt.Printf("%d char is %s\n", num, string(chars))
	return chars
}

func GetNumChar(num int) byte {
	switch num {
	case 0:
		return '0'
	case 1:
		return '1'
	case 2:
		return '2'
	case 3:
		return '3'
	case 4:
		return '4'
	case 5:
		return '5'
	case 6:
		return '6'
	case 7:
		return '7'
	case 8:
		return '8'
	case 9:
		return '9'
	}
	return '0'
}
