package main

func moveZeros(nums []int) {
	i := 0
	j := i + 1
	for i < len(nums) && j < len(nums) {
		if nums[i] != 0 {
			i++
			j++
		} else {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j++
			} else {
				j++
			}
		}
	}
}
