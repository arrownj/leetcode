package main

func removeElements(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	curr, end := 0, len(nums)-1
	for {
		if nums[curr] == val {
			nums[curr], nums[end] = nums[end], nums[curr]
			end--
		} else {
			curr++
		}
		if curr > end {
			break
		}
	}
	return curr
}
