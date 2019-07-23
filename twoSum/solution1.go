package twoSum

func Solution1(nums []int, target int) []int {
	i := 0
	for i < len(nums) {
		j := i + 1
		for j < len(nums) {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
			j++
		}
		i++
	}
	return []int{0, 0}
}
