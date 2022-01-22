package leetcode

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	}

	var j = ln - 1
	for i := 0; i < j; {
		if nums[i] == val && i != j {
			t := nums[j]
			nums[j] = nums[i]
			nums[i] = t
			j--
			continue
		}
		i++
	}

	var k = 0
	for k = 0; k < ln; k++ {
		if nums[k] == val {
			break
		}
	}

	return k
}

// @lc code=end
