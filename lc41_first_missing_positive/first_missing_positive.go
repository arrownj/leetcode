package main

func FirstMissingPositive(nums []int) int {
	for i, num := range nums {
		if num <= 0 {
			nums[i] = len(nums) + 1
		}
	}
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if num <= len(nums) && nums[num-1] > 0 {
			nums[num-1] = 0 - nums[num-1]
		}
	}
	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func AnotherFirstMissingPositive(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}
	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
