package leetcode

/*
 * @lc app=leetcode.cn id=2006 lang=golang
 *
 * [2006] 差的绝对值为 K 的数对数目
 */

// @lc code=start
func countKDifference(nums []int, k int) int {
	result := 0

	var absolute = func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if k == absolute(nums[i], nums[j]) {
				result++
			}
		}
	}

	return result
}

// @lc code=end
