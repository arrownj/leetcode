package lc45_jump_game_2

func Solution(nums []int) int {
	steps := make([]int, len(nums), len(nums))
	for i := 0; i < len(steps); i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j >= len(nums) {
				break
			}
			if steps[i+j] == 0 || steps[i+j] > steps[i]+1 {
				steps[i+j] = steps[i] + 1
			}
		}
	}
	return steps[len(steps)-1]
}
