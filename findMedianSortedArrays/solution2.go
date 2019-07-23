package findMedianSortedArrays

func Solution2(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	min, max, half := 0, m, (m+n+1)/2
	for min <= max {
		i := (min + max) / 2
		j := half - i
		if i-1 >= 0 && nums1[i-1] > nums2[j] {
			max = i - 1
		} else if nums2[j-1] > nums1[i] {
			min = i + 1
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else if nums2[j-1] > nums1[i-1] {
				maxLeft = nums2[j-1]
			} else {
				maxLeft = nums1[i-1]
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			var minRight int
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else if nums1[i] < nums2[j] {
				minRight = nums1[i]
			} else {
				minRight = nums2[j]
			}
			return float64(maxLeft+minRight) / float64(2)
		}
	}
	return float64(0)
}
