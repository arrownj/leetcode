package main

func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target, 0, len(nums)-1)
	last := findLast(nums, target, 0, len(nums)-1)
	return []int{first, last}
}

func findFirst(nums []int, target int, left, right int) int {
	if left > right {
		return -1
	}
	if left == right && nums[left] == target {
		return left
	}
	mid := left + (right-left)/2
	if nums[mid] < target {
		return findFirst(nums, target, mid+1, right)
	} else if nums[mid] > target {
		return findFirst(nums, target, left, mid-1)
	}
	if mid-1 >= 0 && nums[mid-1] == target {
		return findFirst(nums, target, left, mid-1)
	}
	return mid
}

func findLast(nums []int, target int, left, right int) int {
	if left > right {
		return -1
	}
	if left == right && nums[left] == target {
		return left
	}
	mid := left + (right-left)/2
	if nums[mid] > target {
		return findLast(nums, target, left, mid-1)
	} else if nums[mid] < target {
		return findLast(nums, target, mid+1, right)
	}
	if mid+1 < len(nums) && nums[mid+1] == target {
		return findLast(nums, target, mid+1, right)
	}
	return mid
}
