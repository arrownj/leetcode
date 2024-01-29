package main

func rotate(nums []int, k int) []int {
	k = k % len(nums)
	for i := 0; i < k; i++ {
		lastNum := nums[len(nums)-1]
		for j := 0; j < len(nums)-1; j++ {
			nums[len(nums)-j-1] = nums[len(nums)-j-2]
		}
		nums[0] = lastNum
	}
	return nums
}
