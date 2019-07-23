package twoSum

func Solution2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok && j != i {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{0, 0}
}
