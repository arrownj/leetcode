package main

func searchInsert(nums []int, target int) int {
	return searchInPartition(nums, target, 0, len(nums)-1)
}

func searchInPartition(nums []int, target int, left, right int) int {
	if left == right {
		if target <= nums[left] {
			return left
		} else {
			return left + 1
		}
	}

	mid := left + (right-left)/2
	if target <= nums[mid] {
		return searchInPartition(nums, target, left, mid)
	} else {
		return searchInPartition(nums, target, mid+1, right)
	}
}
