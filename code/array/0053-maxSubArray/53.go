package leetcode

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var x = 0
	for i := range nums {
		x += nums[i]
	}

	n1 := nums[:len(nums)-1]
	n2 := nums[1:]

	msa1 := maxSubArray(n1)
	msa2 := maxSubArray(n2)

	if x > msa1 && x > msa2 {
		return x
	}

	if msa1 > msa2 {
		return msa1
	}

	return msa2
}

// @lc code=end
