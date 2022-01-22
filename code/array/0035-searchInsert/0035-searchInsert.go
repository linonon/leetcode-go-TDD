package leetcode

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	// 先判断长度为 1 的情况
	if len(nums) == 1 {
		if nums[0] < target {
			return 1
		} else {
			return 0
		}
	}

	// 判断 target 是不是比第一位小
	if target < nums[0] {
		return 0
	}
	// 判断 target 是不是比最后一个大
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	// 判断中间的数
	var i = 0
	for i = 0; i < len(nums)-1; i++ {
		if nums[i] == target {
			break
		} else if nums[i] < target && target < nums[i+1] {
			i++
			break
		}
	}

	return i
}

// @lc code=end
