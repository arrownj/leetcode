package main

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	last := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[last] {
			last++
			nums[last] = nums[i]
		}
	}
	return last + 1
}
