package findMedianSortedArrays

func Solution1(nums1, nums2 []int) float64 {
	i, j, idx := 0, 0, 0
	nums := make([]int, len(nums1)+len(nums2))
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums[idx] = nums1[i]
			i++
		} else {
			nums[idx] = nums2[j]
			j++
		}
		idx++
	}
	for i < len(nums1) {
		nums[idx] = nums1[i]
		i++
		idx++
	}
	for j < len(nums2) {
		nums[idx] = nums2[j]
		j++
		idx++
	}
	if len(nums)%2 == 1 {
		return float64(nums[(len(nums)-1)/2])
	}
	return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / float64(2)
}
