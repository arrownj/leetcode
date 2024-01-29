package main

func hIndex(citations []int) int {
	nums := make([]int, len(citations)+1)
	for i := 0; i < len(citations); i++ {
		if citations[i] > len(citations) {
			nums[len(citations)]++
		} else {
			nums[citations[i]]++
		}
	}
	h := len(citations)
	for h >= 0 {
		if citations[h] >= h {
			break
		}
		h--
	}
	return h
}
