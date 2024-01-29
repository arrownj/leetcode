package main

func search(nums []int, target int) int {
	return searchBetween(nums, target, 0, len(nums)-1)
}

func searchBetween(nums []int, target int, left, right int) int {
	if left == right && nums[left] != target {
		return -1
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}
	if nums[left] <= nums[mid] {
		if nums[left] <= target && target <= nums[mid] {
			return searchBetween(nums, target, left, mid)
		} else {
			return searchBetween(nums, target, mid+1, right)
		}
	} else {
		if nums[mid+1] <= target && target <= nums[right] {
			return searchBetween(nums, target, mid+1, right)
		} else {
			return searchBetween(nums, target, left, mid)
		}
	}
}
