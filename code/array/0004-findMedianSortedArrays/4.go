package leetcode

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}
	if len(nums1) == 0 {
		if len(nums2)%2 != 1 {
			return float64(nums2[len(nums2)/2]+nums2[len(nums2)/2]-1) / 2
		}
		return float64(nums2[len(nums2)/2])
	}
	if len(nums2) == 0 {
		if len(nums1)%2 != 1 {
			return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2]-1) / 2
		}
		return float64(nums1[len(nums1)/2])
	}

	var ns = make([]int, 0, len(nums1)+len(nums2))
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] <= nums2[j] {
			ns = append(ns, nums1[i])
			i++
		} else {
			ns = append(ns, nums2[j])
			j++
		}
		if i >= len(nums1) {
			ns = append(ns, nums2[j:]...)
		}
		if j >= len(nums2) {
			ns = append(ns, nums1[i:]...)
		}
	}

	if len(ns)%2 != 1 {
		return float64(ns[len(ns)/2]+ns[len(ns)/2-1]) / 2
	}

	return float64(ns[len(ns)/2])

}

// @lc code=end
