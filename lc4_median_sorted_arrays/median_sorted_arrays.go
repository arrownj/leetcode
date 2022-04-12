package main

func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	nums1Len, nums2Len := len(nums1), len(nums2)

	last1, last2 := 0, 0
	left, right := 0, 0
	for last1 < nums1Len && last2 < nums2Len {
		if nums1[last1] < nums2[last2] {
			if (nums1Len+nums2Len)%2 == 1 {
				if last1+last2 == (nums1Len+nums2Len)/2 {
					left, right = nums1[last1], nums1[last1]
					break
				}
			} else {
				if last1+last2 == (nums1Len+nums2Len)/2-1 {
					left = nums1[last1]
				} else if last1+last2 == (nums1Len+nums2Len)/2 {
					right = nums1[last1]
					break
				}
			}
			last1++
		} else {
			if (nums1Len+nums2Len)%2 == 1 {
				if last1+last2 == (nums1Len+nums2Len)/2 {
					left, right = nums2[last2], nums2[last2]
					break
				}
			} else {
				if last1+last2 == (nums1Len+nums2Len)/2-1 {
					left = nums2[last2]
				} else if last1+last2 == (nums1Len+nums2Len)/2 {
					right = nums2[last2]
					break
				}
			}
			last2++
		}
	}
	if last1+last2 < 
	return float64(left+right) / 2
}
