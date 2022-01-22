package leetcode

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return ln
	}

	var j int = 0
	for i := 1; i < ln && j < ln; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}

// @lc code=end
