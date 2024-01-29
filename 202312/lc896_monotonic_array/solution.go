package main

func isMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	if nums[0] == nums[len(nums)-1] {
		for i := 0; i < len(nums); i++ {
			if nums[i] != nums[0] {
				return false
			}
		}
		return true
	}

	if nums[0] > nums[len(nums)-1] {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] < nums[i+1] {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
