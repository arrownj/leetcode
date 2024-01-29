package main

func LongestConsecutiveSequence(nums []int) int {
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}
	longestLength = 0
	for _, num := range nums {
		length := 0
		i := 1
		for {
			if ok, _ := numMap[num+i]; ok {
				length++
			} else {
				break
			}
			i++
		}
		if longestLength < length {
			longestLength = length
		}
	}
	return longestLength
}
