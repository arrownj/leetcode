package min_missing

func Solution(nums []int) int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num-1 != i && num > 0 && num <= len(nums) {
			nums[num-1] = num
		}
	}
	i := 0
	for ; i < len(nums); i++ {
		num := nums[i]
		if num-1 != i {
			break
		}
	}
	return i + 1
}
